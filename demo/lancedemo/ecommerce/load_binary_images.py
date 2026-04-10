#!/usr/bin/env python3
"""
Binary image data loader for H&M Fashion Search with LanceDB
This script downloads images and stores them as base64-encoded binary data in LanceDB.
"""

import pandas as pd
import lancedb
from transformers import CLIPModel, CLIPTokenizer
import numpy as np
import torch
import requests
import base64
from io import BytesIO
from PIL import Image
import time
import os


class Config:
    USE_MINIO = True
    MINIO_ENDPOINT = "http://minio:9000"
    MINIO_ACCESS_KEY = "minioadmin"
    MINIO_SECRET_KEY = "minioadmin123"
    MINIO_BUCKET = "lance"
    MINIO_PATH = "lancedb/"
    MINIO_ALLOW_HTTP = True
    MINIO_REGION = "us-east-1"
    MINIO_FORCE_PATH_STYLE = True
    TABLE_NAME = "hm_mini"
    EMBEDDING_MODEL = "openai/clip-vit-base-patch32"


class CLIPTextEncoder:
    def __init__(self, model_name: str = Config.EMBEDDING_MODEL):
        self.tokenizer = CLIPTokenizer.from_pretrained(model_name)
        self.model = CLIPModel.from_pretrained(model_name)
        self.model.eval()

    def encode_batch(self, texts):
        inputs = self.tokenizer(
            list(texts),
            return_tensors="pt",
            padding=True,
            truncation=True,
            max_length=77,
        )
        with torch.no_grad():
            features = self.model.get_text_features(**inputs)

        # transformers versions may return either a tensor or a model output object.
        if hasattr(features, "pooler_output"):
            features = features.pooler_output
        elif hasattr(features, "last_hidden_state"):
            features = features.last_hidden_state[:, 0, :]

        features = features / torch.norm(features, dim=-1, keepdim=True)
        return features.cpu().numpy().astype(np.float32)


def _build_lancedb_path(use_minio: bool = True):
    if use_minio:
        os.environ["AWS_ACCESS_KEY_ID"] = Config.MINIO_ACCESS_KEY
        os.environ["AWS_SECRET_ACCESS_KEY"] = Config.MINIO_SECRET_KEY
        os.environ["AWS_ENDPOINT_URL_S3"] = Config.MINIO_ENDPOINT
        os.environ["AWS_ENDPOINT_URL"] = Config.MINIO_ENDPOINT
        os.environ["AWS_ALLOW_HTTP"] = str(Config.MINIO_ALLOW_HTTP).lower()
        os.environ["AWS_REGION"] = Config.MINIO_REGION
        os.environ["AWS_S3_FORCE_PATH_STYLE"] = str(Config.MINIO_FORCE_PATH_STYLE).lower()
        os.environ["AWS_VIRTUAL_HOSTED_STYLE_REQUEST"] = "false"
        return f"s3://{Config.MINIO_BUCKET}/{Config.MINIO_PATH}"

    os.environ.pop("AWS_ENDPOINT_URL_S3", None)
    os.environ.pop("AWS_ENDPOINT_URL", None)
    return "./data"

def download_and_encode_image(url, max_size=(400, 400), quality=85):
    """Download an image and encode it as base64"""
    try:
        print(f"Downloading image: {url}")
        
        # Download the image
        response = requests.get(url, timeout=10, headers={
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
        })
        response.raise_for_status()
        
        # Open and process the image
        img = Image.open(BytesIO(response.content))
        
        # Convert to RGB if necessary
        if img.mode in ('RGBA', 'P'):
            img = img.convert('RGB')
        
        # Resize if needed
        img.thumbnail(max_size, Image.Resampling.LANCZOS)
        
        # Save as JPEG to BytesIO
        output = BytesIO()
        img.save(output, format='JPEG', quality=quality, optimize=True)
        output.seek(0)
        
        # Encode as base64
        encoded = base64.b64encode(output.getvalue()).decode('utf-8')
        return f"data:image/jpeg;base64,{encoded}"
        
    except Exception as e:
        print(f"Error downloading/encoding image {url}: {e}")
        # Return a simple colored placeholder as base64
        return create_placeholder_image(max_size, color=(180, 180, 180))

def create_placeholder_image(size=(400, 400), color=(180, 180, 180)):
    """Create a simple placeholder image as base64"""
    try:
        img = Image.new('RGB', size, color)
        output = BytesIO()
        img.save(output, format='JPEG', quality=85)
        output.seek(0)
        encoded = base64.b64encode(output.getvalue()).decode('utf-8')
        return f"data:image/jpeg;base64,{encoded}"
    except Exception as e:
        print(f"Error creating placeholder: {e}")
        return ""

def create_sample_data_with_binary_images():
    """Create sample fashion data with binary image data"""
    
    # Sample image URLs from reliable sources - 55 total images
    image_urls = [
        # Original 10
        "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400&h=400&fit=crop",  # shirt
        "https://images.unsplash.com/photo-1595777457583-95e059d581b8?w=400&h=400&fit=crop",  # dress
        "https://images.unsplash.com/photo-1542272604-787c3835535d?w=400&h=400&fit=crop",   # jeans
        "https://images.unsplash.com/photo-1551698618-1dfe5d97d256?w=400&h=400&fit=crop",   # sneakers
        "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?w=400&h=400&fit=crop",  # dress
        "https://images.unsplash.com/photo-1586790170083-2f9ceadc732d?w=400&h=400&fit=crop",  # jacket
        "https://images.unsplash.com/photo-1617137984095-74e4e5e3613f?w=400&h=400&fit=crop",  # sweater
        "https://images.unsplash.com/photo-1596755094514-f87e34085b2c?w=400&h=400&fit=crop",  # shoes
        "https://images.unsplash.com/photo-1562157873-818bc0726f68?w=400&h=400&fit=crop",   # t-shirt
        "https://images.unsplash.com/photo-1491553895911-0055eca6402d?w=400&h=400&fit=crop",  # shoes
        
        # Additional 45 diverse fashion images
        "https://images.unsplash.com/photo-1544441893-675973e31985?w=400&h=400&fit=crop",   # blazer
        "https://images.unsplash.com/photo-1503342217505-b0a15ec3261c?w=400&h=400&fit=crop",   # sweater
        "https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=400&h=400&fit=crop",   # shoes
        "https://images.unsplash.com/photo-1560243563-062bfc001d68?w=400&h=400&fit=crop",   # dress
        "https://images.unsplash.com/photo-1485218126466-34e6392ec754?w=400&h=400&fit=crop",   # jeans
        "https://images.unsplash.com/photo-1523381210434-271e8be1f52b?w=400&h=400&fit=crop",   # clothing
        "https://images.unsplash.com/photo-1571945153237-4929e783af4a?w=400&h=400&fit=crop",   # boots
        "https://images.unsplash.com/photo-1618932260643-eee4a2f652a6?w=400&h=400&fit=crop",   # coat
        "https://images.unsplash.com/photo-1594633312681-425c7b97ccd1?w=400&h=400&fit=crop",   # sneakers
        "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400&h=400&fit=crop",   # shirt
        "https://images.unsplash.com/photo-1496747611176-843222e1e57c?w=400&h=400&fit=crop",   # dress
        "https://images.unsplash.com/photo-1624378439575-d8705ad7ae80?w=400&h=400&fit=crop",   # pants
        "https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=400&h=400&fit=crop",   # shoes
        "https://images.unsplash.com/photo-1539109136881-3be0616acf4b?w=400&h=400&fit=crop",   # jacket
        "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?w=400&h=400&fit=crop",   # sweater
        "https://images.unsplash.com/photo-1485230895905-ec40ba36b9bc?w=400&h=400&fit=crop",   # blouse
        "https://images.unsplash.com/photo-1529720317453-c8da503f2051?w=400&h=400&fit=crop",   # boots
        "https://images.unsplash.com/photo-1549298916-b41d501d3772?w=400&h=400&fit=crop",   # sneakers
        "https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400&h=400&fit=crop",   # coat
        "https://images.unsplash.com/photo-1516762689617-e1cffcef479d?w=400&h=400&fit=crop",   # suit
        "https://images.unsplash.com/photo-1534670007418-fbb7f6cf32c3?w=400&h=400&fit=crop",   # sandals
        "https://images.unsplash.com/photo-1515886657613-9f3515b0c78f?w=400&h=400&fit=crop",   # fashion
        "https://images.unsplash.com/photo-1591047139829-d91aecb6caea?w=400&h=400&fit=crop",   # pants
        "https://images.unsplash.com/photo-1553062407-98eeb64c6a62?w=400&h=400&fit=crop",   # men shoes
        "https://images.unsplash.com/photo-1572804013309-59a88b7e92f1?w=400&h=400&fit=crop",   # dress
        "https://images.unsplash.com/photo-1549298916-b41d501d3772?w=400&h=400&fit=crop",   # boots
        "https://images.unsplash.com/photo-1445205170230-053b83016050?w=400&h=400&fit=crop",   # shirt
        "https://images.unsplash.com/photo-1485230895905-ec40ba36b9bc?w=400&h=400&fit=crop",   # blouse
        "https://images.unsplash.com/photo-1542272604-787c3835535d?w=400&h=400&fit=crop",   # jeans
        "https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=400&h=400&fit=crop",   # shoes
        "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?w=400&h=400&fit=crop",   # sweater
        "https://images.unsplash.com/photo-1485462537746-965f33f7f6a7?w=400&h=400&fit=crop",   # dress
        "https://images.unsplash.com/photo-1539109136881-3be0616acf4b?w=400&h=400&fit=crop",   # jacket
        "https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=400&h=400&fit=crop",   # shoes
        "https://images.unsplash.com/photo-1445205170230-053b83016050?w=400&h=400&fit=crop",   # outfit
        "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400&h=400&fit=crop",   # shirt
        "https://images.unsplash.com/photo-1624378439575-d8705ad7ae80?w=400&h=400&fit=crop",   # pants
        "https://images.unsplash.com/photo-1485230895905-ec40ba36b9bc?w=400&h=400&fit=crop",   # blouse
        "https://images.unsplash.com/photo-1529720317453-c8da503f2051?w=400&h=400&fit=crop",   # accessories
        "https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=400&h=400&fit=crop",   # shoes
        "https://images.unsplash.com/photo-1485230895905-ec40ba36b9bc?w=400&h=400&fit=crop",   # fashion
        "https://images.unsplash.com/photo-1556821840-3a9c6494e5bd?w=400&h=400&fit=crop",   # clothing
        "https://images.unsplash.com/photo-1445205170230-053b83016050?w=400&h=400&fit=crop",   # shirt
        "https://images.unsplash.com/photo-1610030469983-98e550d6193c?w=400&h=400&fit=crop",   # sneakers
        "https://images.unsplash.com/photo-1578979879663-4ba3abd39d32?w=400&h=400&fit=crop",   # dress
        "https://images.unsplash.com/photo-1566479179817-c04deb4b0e9a?w=400&h=400&fit=crop",   # fashion
    ]
    
    sample_data = [
        {
            "image_url": "binary_stored",  # Indicate that image is stored as binary
            "prod_name": "Classic Cotton Shirt",
            "detail_desc": "A comfortable and stylish cotton shirt perfect for casual and formal occasions. Features a modern fit with a classic collar design.",
            "product_type_name": "Shirt",
            "index_group_name": "Menswear",
            "price": 29.99,
            "article_id": "MS001",
            "available": True,
            "color": "Blue",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Elegant Summer Dress",
            "detail_desc": "Light and breezy summer dress with floral patterns. Made from breathable fabric perfect for warm weather.",
            "product_type_name": "Dress",
            "index_group_name": "Ladieswear",
            "price": 45.50,
            "article_id": "LD001",
            "available": True,
            "color": "Floral",
            "size": "S"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Slim Fit Jeans",
            "detail_desc": "Modern slim-fit jeans with stretch fabric for comfort and style. Perfect for everyday wear.",
            "product_type_name": "Jeans",
            "index_group_name": "Menswear",
            "price": 59.99,
            "article_id": "MJ001",
            "available": True,
            "color": "Dark Blue",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Athletic Sneakers",
            "detail_desc": "Comfortable athletic sneakers with advanced cushioning technology. Perfect for sports and casual wear.",
            "product_type_name": "Shoes",
            "index_group_name": "Sport",
            "price": 89.99,
            "article_id": "SS001",
            "available": True,
            "color": "White",
            "size": "42"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Business Dress",
            "detail_desc": "Professional business dress suitable for office wear. Features a tailored fit and premium fabric.",
            "product_type_name": "Dress",
            "index_group_name": "Ladieswear",
            "price": 79.99,
            "article_id": "LD002",
            "available": True,
            "color": "Black",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Winter Jacket",
            "detail_desc": "Warm winter jacket with insulation and water-resistant coating. Perfect for cold weather.",
            "product_type_name": "Jacket",
            "index_group_name": "Menswear",
            "price": 129.99,
            "article_id": "MJ002",
            "available": True,
            "color": "Navy",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Cozy Sweater",
            "detail_desc": "Soft and warm sweater made from premium wool blend. Perfect for layering in cooler weather.",
            "product_type_name": "Sweater",
            "index_group_name": "Ladieswear",
            "price": 55.00,
            "article_id": "LS001",
            "available": True,
            "color": "Cream",
            "size": "S"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Casual Loafers",
            "detail_desc": "Elegant casual loafers made from genuine leather. Perfect for both casual and semi-formal occasions.",
            "product_type_name": "Shoes",
            "index_group_name": "Menswear",
            "price": 95.00,
            "article_id": "MS002",
            "available": True,
            "color": "Brown",
            "size": "43"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Basic T-Shirt",
            "detail_desc": "High-quality basic t-shirt made from 100% organic cotton. A wardrobe essential in various colors.",
            "product_type_name": "T-shirt",
            "index_group_name": "Divided",
            "price": 12.99,
            "article_id": "DT001",
            "available": True,
            "color": "White",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Running Shoes",
            "detail_desc": "High-performance running shoes with advanced sole technology and breathable upper material.",
            "product_type_name": "Shoes",
            "index_group_name": "Sport",
            "price": 119.99,
            "article_id": "SS002",
            "available": True,
            "color": "Black/Red",
            "size": "41"
        },
        
        # Additional 45 items
        {
            "image_url": "binary_stored",
            "prod_name": "Classic Blazer",
            "detail_desc": "Sophisticated blazer perfect for business meetings and formal occasions. Tailored fit with premium fabric.",
            "product_type_name": "Blazer",
            "index_group_name": "Menswear",
            "price": 149.99,
            "article_id": "MB001",
            "available": True,
            "color": "Navy",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Wool Pullover",
            "detail_desc": "Cozy wool pullover sweater with ribbed cuffs and hem. Perfect for layering in cool weather.",
            "product_type_name": "Sweater",
            "index_group_name": "Ladieswear",
            "price": 68.00,
            "article_id": "LW001",
            "available": True,
            "color": "Gray",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Oxford Shoes",
            "detail_desc": "Classic oxford dress shoes made from genuine leather. Perfect for formal and business attire.",
            "product_type_name": "Shoes",
            "index_group_name": "Menswear",
            "price": 125.00,
            "article_id": "MS003",
            "available": True,
            "color": "Black",
            "size": "42"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Midi Dress",
            "detail_desc": "Elegant midi dress with floral print. Features a flattering A-line silhouette and three-quarter sleeves.",
            "product_type_name": "Dress",
            "index_group_name": "Ladieswear",
            "price": 89.99,
            "article_id": "LD003",
            "available": True,
            "color": "Floral",
            "size": "S"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Distressed Jeans",
            "detail_desc": "Trendy distressed jeans with a relaxed fit. Features stylish rips and vintage wash.",
            "product_type_name": "Jeans",
            "index_group_name": "Divided",
            "price": 49.99,
            "article_id": "DJ001",
            "available": True,
            "color": "Light Blue",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Casual Cardigan",
            "detail_desc": "Soft knit cardigan with button closure. Perfect for layering over dresses or with jeans.",
            "product_type_name": "Cardigan",
            "index_group_name": "Ladieswear",
            "price": 55.00,
            "article_id": "LC001",
            "available": True,
            "color": "Beige",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Combat Boots",
            "detail_desc": "Edgy combat boots with lace-up design and platform sole. Perfect for adding attitude to any outfit.",
            "product_type_name": "Boots",
            "index_group_name": "Divided",
            "price": 79.99,
            "article_id": "DB001",
            "available": True,
            "color": "Black",
            "size": "38"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Peacoat",
            "detail_desc": "Classic double-breasted peacoat in wool blend. Timeless design perfect for cold weather.",
            "product_type_name": "Coat",
            "index_group_name": "Menswear",
            "price": 189.99,
            "article_id": "MC001",
            "available": True,
            "color": "Charcoal",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Sports Bra",
            "detail_desc": "High-support sports bra with moisture-wicking fabric. Perfect for intense workouts and running.",
            "product_type_name": "Sports Bra",
            "index_group_name": "Sport",
            "price": 29.99,
            "article_id": "SB001",
            "available": True,
            "color": "Pink",
            "size": "S"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Polo Shirt",
            "detail_desc": "Classic polo shirt in cotton pique. Features a traditional collar and three-button placket.",
            "product_type_name": "Polo",
            "index_group_name": "Menswear",
            "price": 35.00,
            "article_id": "MP001",
            "available": True,
            "color": "White",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Maxi Dress",
            "detail_desc": "Flowing maxi dress with bohemian print. Perfect for summer festivals and beach vacations.",
            "product_type_name": "Dress",
            "index_group_name": "Ladieswear",
            "price": 75.00,
            "article_id": "LD004",
            "available": True,
            "color": "Multi",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Chino Pants",
            "detail_desc": "Versatile chino pants in stretch cotton. Perfect for both casual and smart-casual occasions.",
            "product_type_name": "Pants",
            "index_group_name": "Menswear",
            "price": 59.99,
            "article_id": "MP002",
            "available": True,
            "color": "Khaki",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Ballet Flats",
            "detail_desc": "Comfortable ballet flats with padded insole. Classic design that goes with everything.",
            "product_type_name": "Flats",
            "index_group_name": "Ladieswear",
            "price": 45.00,
            "article_id": "LF001",
            "available": True,
            "color": "Black",
            "size": "37"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Leather Jacket",
            "detail_desc": "Genuine leather jacket with asymmetric zip and quilted shoulders. Timeless biker style.",
            "product_type_name": "Jacket",
            "index_group_name": "Divided",
            "price": 199.99,
            "article_id": "DJ002",
            "available": True,
            "color": "Black",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Cashmere Scarf",
            "detail_desc": "Luxurious cashmere scarf in solid color. Soft texture and elegant drape for any outfit.",
            "product_type_name": "Scarf",
            "index_group_name": "Ladieswear",
            "price": 89.00,
            "article_id": "LS002",
            "available": True,
            "color": "Cream",
            "size": "One Size"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Silk Blouse",
            "detail_desc": "Elegant silk blouse with button-down design. Perfect for office wear or special occasions.",
            "product_type_name": "Blouse",
            "index_group_name": "Ladieswear",
            "price": 95.00,
            "article_id": "LB001",
            "available": True,
            "color": "Ivory",
            "size": "S"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Chelsea Boots",
            "detail_desc": "Classic Chelsea boots in suede with elastic side panels. Versatile style for any wardrobe.",
            "product_type_name": "Boots",
            "index_group_name": "Menswear",
            "price": 135.00,
            "article_id": "MB002",
            "available": True,
            "color": "Brown",
            "size": "43"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Yoga Leggings",
            "detail_desc": "High-waisted yoga leggings with four-way stretch. Perfect for yoga, pilates, and everyday wear.",
            "product_type_name": "Leggings",
            "index_group_name": "Sport",
            "price": 39.99,
            "article_id": "SL001",
            "available": True,
            "color": "Black",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Denim Jacket",
            "detail_desc": "Classic denim jacket with vintage wash. Features chest pockets and button closure.",
            "product_type_name": "Jacket",
            "index_group_name": "Divided",
            "price": 69.99,
            "article_id": "DJ003",
            "available": True,
            "color": "Blue",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Three-Piece Suit",
            "detail_desc": "Formal three-piece suit in wool blend. Includes jacket, vest, and trousers for special occasions.",
            "product_type_name": "Suit",
            "index_group_name": "Menswear",
            "price": 299.99,
            "article_id": "MS004",
            "available": True,
            "color": "Navy",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Platform Sandals",
            "detail_desc": "Trendy platform sandals with ankle strap. Perfect for adding height and style to summer outfits.",
            "product_type_name": "Sandals",
            "index_group_name": "Ladieswear",
            "price": 65.00,
            "article_id": "LS003",
            "available": True,
            "color": "Tan",
            "size": "39"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Crop Top",
            "detail_desc": "Trendy crop top with ribbed texture. Perfect for layering or wearing alone in warm weather.",
            "product_type_name": "Top",
            "index_group_name": "Divided",
            "price": 19.99,
            "article_id": "DT002",
            "available": True,
            "color": "White",
            "size": "S"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Cargo Pants",
            "detail_desc": "Utility cargo pants with multiple pockets. Perfect for outdoor activities and casual wear.",
            "product_type_name": "Pants",
            "index_group_name": "Menswear",
            "price": 65.00,
            "article_id": "MP003",
            "available": True,
            "color": "Olive",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "High Heels",
            "detail_desc": "Classic pointed-toe high heels. Elegant design perfect for evening wear and special occasions.",
            "product_type_name": "Heels",
            "index_group_name": "Ladieswear",
            "price": 85.00,
            "article_id": "LH001",
            "available": True,
            "color": "Black",
            "size": "38"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Hoodie",
            "detail_desc": "Comfortable cotton hoodie with kangaroo pocket. Perfect for casual wear and lounging.",
            "product_type_name": "Hoodie",
            "index_group_name": "Divided",
            "price": 39.99,
            "article_id": "DH001",
            "available": True,
            "color": "Gray",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Button-Down Shirt",
            "detail_desc": "Crisp button-down shirt in premium cotton. Essential piece for professional and casual wardrobes.",
            "product_type_name": "Shirt",
            "index_group_name": "Menswear",
            "price": 49.99,
            "article_id": "MS005",
            "available": True,
            "color": "Light Blue",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Pleated Skirt",
            "detail_desc": "Classic pleated midi skirt with elastic waistband. Versatile piece that works for office or weekend.",
            "product_type_name": "Skirt",
            "index_group_name": "Ladieswear",
            "price": 52.00,
            "article_id": "LS004",
            "available": True,
            "color": "Navy",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Athletic Shorts",
            "detail_desc": "Lightweight athletic shorts with moisture-wicking fabric and side pockets for active lifestyle.",
            "product_type_name": "Shorts",
            "index_group_name": "Sport",
            "price": 24.99,
            "article_id": "SS003",
            "available": True,
            "color": "Navy",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Wrap Dress",
            "detail_desc": "Flattering wrap dress with tie waist. Classic silhouette that suits all body types.",
            "product_type_name": "Dress",
            "index_group_name": "Ladieswear",
            "price": 69.99,
            "article_id": "LD005",
            "available": True,
            "color": "Red",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Bomber Jacket",
            "detail_desc": "Trendy bomber jacket with ribbed cuffs and hem. Perfect for adding a casual edge to any outfit.",
            "product_type_name": "Jacket",
            "index_group_name": "Divided",
            "price": 79.99,
            "article_id": "DJ004",
            "available": True,
            "color": "Black",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Loafers",
            "detail_desc": "Comfortable leather loafers with classic penny design. Perfect for both casual and business casual looks.",
            "product_type_name": "Loafers",
            "index_group_name": "Menswear",
            "price": 89.99,
            "article_id": "ML001",
            "available": True,
            "color": "Brown",
            "size": "42"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Tank Top",
            "detail_desc": "Basic cotton tank top with slim fit. Essential layering piece for any wardrobe.",
            "product_type_name": "Tank",
            "index_group_name": "Divided",
            "price": 14.99,
            "article_id": "DT003",
            "available": True,
            "color": "Black",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Wide-Leg Pants",
            "detail_desc": "Flowy wide-leg pants with high waist. Comfortable and stylish for both casual and dressy occasions.",
            "product_type_name": "Pants",
            "index_group_name": "Ladieswear",
            "price": 58.00,
            "article_id": "LP001",
            "available": True,
            "color": "Black",
            "size": "S"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Slip-On Sneakers",
            "detail_desc": "Comfortable slip-on sneakers with canvas upper. Easy to wear for everyday casual activities.",
            "product_type_name": "Sneakers",
            "index_group_name": "Divided",
            "price": 45.00,
            "article_id": "DS001",
            "available": True,
            "color": "White",
            "size": "40"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Turtleneck Sweater",
            "detail_desc": "Cozy turtleneck sweater in merino wool blend. Perfect for layering during colder months.",
            "product_type_name": "Sweater",
            "index_group_name": "Menswear",
            "price": 75.00,
            "article_id": "MS006",
            "available": True,
            "color": "Black",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Pencil Skirt",
            "detail_desc": "Classic pencil skirt with back slit. Professional piece perfect for office wear and formal occasions.",
            "product_type_name": "Skirt",
            "index_group_name": "Ladieswear",
            "price": 48.00,
            "article_id": "LS005",
            "available": True,
            "color": "Gray",
            "size": "S"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Training Shoes",
            "detail_desc": "Cross-training shoes with superior grip and support. Perfect for gym workouts and fitness activities.",
            "product_type_name": "Training Shoes",
            "index_group_name": "Sport",
            "price": 95.00,
            "article_id": "ST001",
            "available": True,
            "color": "Gray/Blue",
            "size": "41"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Flannel Shirt",
            "detail_desc": "Soft flannel shirt with classic plaid pattern. Perfect for layering and casual weekend wear.",
            "product_type_name": "Shirt",
            "index_group_name": "Menswear",
            "price": 42.00,
            "article_id": "MS007",
            "available": True,
            "color": "Plaid",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Bodysuit",
            "detail_desc": "Sleek bodysuit with snap closure. Perfect for tucking into skirts and pants for a polished look.",
            "product_type_name": "Bodysuit",
            "index_group_name": "Ladieswear",
            "price": 35.00,
            "article_id": "LB002",
            "available": True,
            "color": "Black",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Swim Shorts",
            "detail_desc": "Quick-dry swim shorts with mesh lining and elastic waistband. Perfect for beach and pool activities.",
            "product_type_name": "Swimwear",
            "index_group_name": "Sport",
            "price": 29.99,
            "article_id": "SS004",
            "available": True,
            "color": "Blue",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Puffer Vest",
            "detail_desc": "Lightweight puffer vest with down filling. Perfect layering piece for transitional weather.",
            "product_type_name": "Vest",
            "index_group_name": "Menswear",
            "price": 65.00,
            "article_id": "MV001",
            "available": True,
            "color": "Navy",
            "size": "L"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Midi Skirt",
            "detail_desc": "A-line midi skirt with side zip closure. Versatile piece that works for both casual and formal occasions.",
            "product_type_name": "Skirt",
            "index_group_name": "Ladieswear",
            "price": 45.00,
            "article_id": "LS006",
            "available": True,
            "color": "Burgundy",
            "size": "M"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Canvas Sneakers",
            "detail_desc": "Classic canvas sneakers with rubber sole. Timeless design perfect for casual everyday wear.",
            "product_type_name": "Sneakers",
            "index_group_name": "Divided",
            "price": 35.00,
            "article_id": "DS002",
            "available": True,
            "color": "Red",
            "size": "39"
        },
        {
            "image_url": "binary_stored",
            "prod_name": "Track Jacket",
            "detail_desc": "Sporty track jacket with zip closure and side pockets. Perfect for workouts and athleisure looks.",
            "product_type_name": "Track Jacket",
            "index_group_name": "Sport",
            "price": 55.00,
            "article_id": "ST002",
            "available": True,
            "color": "Black/White",
            "size": "M"
        }
    ]
    
    # Download and encode images
    print("Downloading and encoding images...")
    for i, item in enumerate(sample_data):
        if i < len(image_urls):
            item["image_data"] = download_and_encode_image(image_urls[i])
        else:
            # Create placeholder for items without specific image URLs
            item["image_data"] = create_placeholder_image()
        
        # Small delay to be respectful to image servers
        time.sleep(0.5)
    
    return sample_data

def main():
    db_path = _build_lancedb_path(Config.USE_MINIO)
    print(f"Connecting to LanceDB at {db_path}...")
    db = lancedb.connect(db_path)
    table_name = Config.TABLE_NAME
    
    print(f"Creating sample data with binary images...")
    sample_data = create_sample_data_with_binary_images()
    
    # Create DataFrame
    df = pd.DataFrame(sample_data)
    
    print("Loading CLIP text encoder...")
    model = CLIPTextEncoder(Config.EMBEDDING_MODEL)

    # Generate embeddings from product descriptions
    print("Generating embeddings...")
    descriptions = df['detail_desc'].fillna("").tolist()
    embeddings = model.encode_batch(descriptions)
    
    # Add embeddings to the DataFrame
    df['vector'] = embeddings.tolist()
    
    # Append mode: open existing table and add rows, or create table when missing.
    try:
        table = db.open_table(table_name)
        print(f"Opened existing table: {table_name}")
        print(f"Appending {len(df)} items...")
        table.add(df)
    except Exception as e:
        message = str(e).lower()
        table_missing = (
            "not found" in message
            or "does not exist" in message
            or "no such table" in message
            or "table" in message and "missing" in message
        )
        if table_missing:
            print(f"Creating table {table_name} with {len(df)} items...")
            table = db.create_table(table_name, df)
        else:
            raise RuntimeError(f"Failed to open table '{table_name}' on LanceDB: {e}")
    
    print(f"✅ Successfully upserted {len(df)} items into table '{table_name}' with binary image data!")
    print("\nSample article IDs for testing:")
    for i, row in df.iterrows():
        print(f"  - {row['article_id']}: {row['prod_name']}")
    
    print(f"\n🖼️ Images are now stored as binary data in the database!")
    print(f"📡 You can access images via: http://localhost:8000/image/<article_id>")
    print(f"🌐 Example: http://localhost:8000/image/MS001")

if __name__ == "__main__":
    main()
