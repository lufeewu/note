from fastapi import FastAPI, HTTPException, Query
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles
from fastapi.responses import Response
import lancedb
from transformers import CLIPModel, CLIPTokenizer
import torch
from typing import List, Optional
from pydantic import BaseModel
from urllib.parse import unquote
import numpy as np
import pandas as pd
import base64
import os
import pyarrow as pa


class CLIPTextEncoder:
    """Lightweight CLIP text encoder using transformers directly."""

    def __init__(self, model_name: str = "openai/clip-vit-base-patch32"):
        self.tokenizer = CLIPTokenizer.from_pretrained(model_name)
        self.model = CLIPModel.from_pretrained(model_name)
        self.model.eval()

    def encode(self, text: str) -> np.ndarray:
        inputs = self.tokenizer(
            text, return_tensors="pt", padding=True, truncation=True, max_length=77
        )
        with torch.no_grad():
            features = self.model.get_text_features(**inputs)

        # transformers versions may return either a tensor or a model output object.
        if hasattr(features, "pooler_output"):
            features = features.pooler_output
        elif hasattr(features, "last_hidden_state"):
            features = features.last_hidden_state[:, 0, :]

        features = features / torch.norm(features, dim=-1, keepdim=True)
        return features.squeeze().cpu().numpy().astype(np.float32)

# Configuration
class Config:
    # 本地 LanceDB 路径
    LANCEDB_PATH = "./data"
    # MinIO/S3 配置（如需使用 MinIO，填写下方配置并设置 USE_MINIO=True）
    USE_MINIO = True
    MINIO_ENDPOINT = "http://minio:9000"
    MINIO_ACCESS_KEY = "minioadmin"
    MINIO_SECRET_KEY = "minioadmin123"
    MINIO_BUCKET = "lance"
    MINIO_PATH = "lancedb/"
    MINIO_ALLOW_HTTP = True
    MINIO_REGION = "us-east-1"
    MINIO_FORCE_PATH_STYLE = True
    # 表名、模型等
    TABLE_NAME = "hm_mini"
    EMBEDDING_MODEL = "openai/clip-vit-base-patch32"
    GROUP_ORDER = ["Menswear", "Ladieswear", "Divided", "Baby/Children", "Sport"]

class SearchResult(BaseModel):
    image_url: str
    prod_name: str
    detail_desc: str
    product_type_name: str
    index_group_name: str
    price: float = 0.0
    article_id: str = ""
    available: bool = True
    color: str = ""
    size: str = ""

    class Config:
        from_attributes = True

class LanceDBService:
    def __init__(self):
        self.encoder = CLIPTextEncoder(Config.EMBEDDING_MODEL)
        self.table = None
        self.db = self._connect_db_minio_only()

    def _set_storage_env(self, use_minio: bool):
        if use_minio:
            os.environ["AWS_ACCESS_KEY_ID"] = Config.MINIO_ACCESS_KEY
            os.environ["AWS_SECRET_ACCESS_KEY"] = Config.MINIO_SECRET_KEY
            os.environ["AWS_ENDPOINT_URL_S3"] = Config.MINIO_ENDPOINT
            os.environ["AWS_ENDPOINT_URL"] = Config.MINIO_ENDPOINT
            os.environ["AWS_ALLOW_HTTP"] = str(Config.MINIO_ALLOW_HTTP).lower()
            os.environ["AWS_REGION"] = Config.MINIO_REGION
            os.environ["AWS_S3_FORCE_PATH_STYLE"] = str(Config.MINIO_FORCE_PATH_STYLE).lower()
            os.environ["AWS_VIRTUAL_HOSTED_STYLE_REQUEST"] = "false"
        else:
            # Ensure local mode does not accidentally reuse external AWS envs.
            os.environ.pop("AWS_ENDPOINT_URL_S3", None)
            os.environ.pop("AWS_ENDPOINT_URL", None)

    def _connect_db_minio_only(self):
        if not Config.USE_MINIO:
            raise RuntimeError("MinIO-only mode enabled: set Config.USE_MINIO = True")

        lancedb_path = f"s3://{Config.MINIO_BUCKET}/{Config.MINIO_PATH}"
        self._set_storage_env(True)
        print(f"Connecting to LanceDB (MinIO/S3) at: {lancedb_path}")

        try:
            self.db = lancedb.connect(lancedb_path)
            print(f"Connected to LanceDB via MinIO successfully: {lancedb_path}")
            self._ensure_table_exists()
            return self.db
        except Exception as e:
            raise RuntimeError(f"Failed to initialize LanceDB via MinIO only: {e}")

    def _get_table_schema(self) -> pa.Schema:
        return pa.schema([
            pa.field("image_url", pa.string()),
            pa.field("image_data", pa.string()),
            pa.field("prod_name", pa.string()),
            pa.field("detail_desc", pa.string()),
            pa.field("product_type_name", pa.string()),
            pa.field("index_group_name", pa.string()),
            pa.field("price", pa.float64()),
            pa.field("article_id", pa.string()),
            pa.field("available", pa.bool_()),
            pa.field("color", pa.string()),
            pa.field("size", pa.string()),
            pa.field("vector", pa.list_(pa.float32(), 512)),
        ])

    def _ensure_table_exists(self):
        """Ensure the table exists, create if it doesn't"""
        try:
            self.table = self.db.open_table(Config.TABLE_NAME)
            print(f"Connected to LanceDB table '{Config.TABLE_NAME}' successfully.")
        except Exception as e:
            message = str(e).lower()
            table_missing = (
                "not found" in message
                or "does not exist" in message
                or "no such table" in message
                or "table" in message and "missing" in message
            )
            if table_missing:
                self.table = self.db.create_table(
                    Config.TABLE_NAME,
                    schema=self._get_table_schema(),
                )
                print(f"Created LanceDB table '{Config.TABLE_NAME}' successfully.")
            else:
                raise RuntimeError(f"Failed to open LanceDB table '{Config.TABLE_NAME}': {e}")

        if self.table is None:
            raise RuntimeError(f"LanceDB table '{Config.TABLE_NAME}' is not initialized")

    def create_filter(self, groups: List[str] = None, items: List[str] = None) -> str:
        """Create LanceDB filter string"""
        conditions = []
        if groups:
            group_condition = " OR ".join([f"index_group_name = '{group}'" for group in groups])
            conditions.append(f"({group_condition})")
        if items:
            item_condition = " OR ".join([f"product_type_name = '{item}'" for item in items])
            conditions.append(f"({item_condition})")
        
        if conditions:
            return " AND ".join(conditions)
        return None

    async def search(self, query: str, groups: List[str], items: List[str], 
                    limit: int, offset: int) -> List[SearchResult]:
        try:
            filter_condition = self.create_filter(groups, items)
            
            if not query:
                # No query, just filter and paginate
                if filter_condition:
                    results = self.table.search().where(filter_condition).limit(limit).to_pandas()
                else:
                    results = self.table.search().limit(limit).to_pandas()
            else:
                # Semantic search with query
                query_vector = self.encoder.encode(query).tolist()
                
                search_query = self.table.search(query_vector)
                if filter_condition:
                    search_query = search_query.where(filter_condition)
                
                results = search_query.limit(limit).to_pandas()

            if results.empty:
                return []

            # Convert DataFrame to list of SearchResult objects
            search_results = []
            for _, row in results.iterrows():
                search_results.append(SearchResult(
                    image_url=row.get('image_url', ''),
                    prod_name=row.get('prod_name', 'Unknown Product'),
                    detail_desc=row.get('detail_desc', 'No description available'),
                    product_type_name=row.get('product_type_name', ''),
                    index_group_name=row.get('index_group_name', ''),
                    price=float(row.get('price', 0.0)),
                    article_id=row.get('article_id', ''),
                    available=row.get('available', True),
                    color=row.get('color', ''),
                    size=row.get('size', '')
                ))
            
            return search_results

        except Exception as e:
            raise HTTPException(
                status_code=500, 
                detail=f"Search error: {str(e)}"
            )

    async def get_groups(self) -> List[str]:
        try:
            # Get unique groups from the table
            groups_df = self.table.search().select(["index_group_name"]).to_pandas()
            
            if groups_df.empty:
                return []
            
            # Get unique values and filter out None/NaN
            groups = set()
            for group in groups_df['index_group_name'].dropna().unique():
                if group and str(group).strip():
                    groups.add(str(group).strip())
            
            # Sort according to GROUP_ORDER preference
            ordered_groups = [g for g in Config.GROUP_ORDER if g in groups]
            remaining_groups = sorted([g for g in groups if g not in Config.GROUP_ORDER])
            
            return ordered_groups + remaining_groups
            
        except Exception as e:
            raise HTTPException(
                status_code=500, 
                detail=f"Failed to fetch groups: {str(e)}"
            )

# Initialize FastAPI and services
app = FastAPI(title="H&M Fashion Search API")
lancedb_service = LanceDBService()

# Configure CORS and static files
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
app.mount("/static", StaticFiles(directory="static"), name="static")

@app.get("/search", response_model=List[SearchResult])
async def search_fashion_items(
    query: str = "", 
    group: List[str] = Query(default=[]),
    item: List[str] = Query(default=[]),
    limit: int = 20
):
    """Search for fashion items using semantic search and/or filters."""
    query = unquote(query.strip())
    groups = [unquote(g.strip()) for g in group]
    items = [unquote(i.strip()) for i in item]
    return await lancedb_service.search(query, groups, items, limit, 0)

@app.get("/groups", response_model=List[str])
async def get_groups():
    """Get list of unique index group names in specified order."""
    return await lancedb_service.get_groups()

@app.get("/image/{article_id}")
async def get_image(article_id: str):
    """Serve binary image data from the database."""
    try:
        # Search for the item by article_id
        results = lancedb_service.table.search().where(f"article_id = '{article_id}'").limit(1).to_pandas()
        
        if results.empty:
            raise HTTPException(status_code=404, detail="Image not found")
        
        image_data = results.iloc[0]['image_data']
        
        # Decode base64 image data
        try:
            # Remove data URL prefix if present (e.g., "data:image/jpeg;base64,")
            if image_data.startswith('data:'):
                image_data = image_data.split(',', 1)[1]
            
            image_bytes = base64.b64decode(image_data)
            
            # Determine content type based on image format
            if image_bytes.startswith(b'\xff\xd8\xff'):
                content_type = "image/jpeg"
            elif image_bytes.startswith(b'\x89PNG'):
                content_type = "image/png"
            elif image_bytes.startswith(b'GIF'):
                content_type = "image/gif"
            elif image_bytes.startswith(b'WEBP'):
                content_type = "image/webp"
            else:
                content_type = "image/jpeg"  # Default fallback
            
            return Response(content=image_bytes, media_type=content_type)
            
        except Exception as e:
            raise HTTPException(status_code=500, detail=f"Error decoding image: {str(e)}")
            
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error retrieving image: {str(e)}")
