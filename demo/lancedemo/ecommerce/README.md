# H&M Fashion Search with LanceDB

A modern fashion search application built with FastAPI and LanceDB, featuring semantic search capabilities for H&M fashion items.

## Features

- **Semantic Search**: Find fashion items using natural language queries
- **Filtering**: Filter by product groups and item types
- **Real-time Results**: Fast search results with pagination
- **Modern UI**: Clean, responsive interface with product details modal
- **Vector Database**: Powered by LanceDB for efficient similarity search

## Technology Stack

- **Backend**: FastAPI (Python)
- **Vector Database**: LanceDB
- **Embeddings**: Sentence Transformers (CLIP ViT-B-32)
- **Frontend**: HTML, CSS, JavaScript
- **AI Models**: Hugging Face Transformers

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd lancedb-ecommerce
```

2. Install dependencies:
```bash
pip install -r requirements.txt
```

3. Create the data directory:
```bash
mkdir -p data
```

## Usage

### Starting the Application

1. Run the FastAPI server:
```bash
uvicorn app:app --reload --host 0.0.0.0 --port 8000
```

2. Open your browser and navigate to `http://localhost:8000/static/index.html`

### Adding Data to LanceDB

The application will automatically create an empty table when it starts. To populate it with fashion data:

1. Prepare your data in a format compatible with the schema:
```python
import pandas as pd
import lancedb
from sentence_transformers import SentenceTransformer

# Connect to the database
db = lancedb.connect("./data")
table = db.open_table("h&m-mini")

# Load your fashion data
data = pd.read_csv("your_fashion_data.csv")

# Generate embeddings for text fields (e.g., product descriptions)
encoder = SentenceTransformer("clip-ViT-B-32")
data['vector'] = data['detail_desc'].apply(lambda x: encoder.encode(x).tolist())

# Insert data into the table
table.add(data)
```

2. The expected schema includes:
   - `image_url`: URL to product image
   - `prod_name`: Product name
   - `detail_desc`: Product description
   - `product_type_name`: Type of product
   - `index_group_name`: Product category group
   - `price`: Product price
   - `article_id`: Unique article identifier
   - `available`: Availability status
   - `color`: Product color
   - `size`: Product size
   - `vector`: 512-dimensional embedding vector

## API Endpoints

- `GET /search`: Search for fashion items with optional filters
- `GET /groups`: Get available product groups
- `GET /static/index.html`: Main application interface

## Configuration

Edit the `Config` class in `app.py` to customize:
- Database path
- Table name
- Embedding model
- Product group ordering

## Differences from Qdrant Version

This LanceDB version offers several advantages:
- **Local Storage**: No need for external database server
- **Simplified Setup**: Automatic table creation and management
- **Better Performance**: Optimized for local development and testing
- **Easier Deployment**: Single file database that can be easily shared

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the MIT License.

