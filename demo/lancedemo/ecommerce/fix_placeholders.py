#!/usr/bin/env python3
"""
Fix placeholder images for specific items in the LanceDB database
"""

import pandas as pd
import lancedb
import requests
import base64
from io import BytesIO
from PIL import Image
import time

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
        return None

def fix_placeholder_images():
    """Fix specific items that have placeholder images"""
    
    # Connect to LanceDB
    db = lancedb.connect("./data")
    table_name = "hm_mini"
    table = db.open_table(table_name)
    
    # Items to fix with new image URLs
    fixes = {
        "MP001": "https://images.unsplash.com/photo-1562157873-818bc0726f68?w=400&h=400&fit=crop",  # polo shirt
        "LS005": "https://images.unsplash.com/photo-1594633312681-425c7b97ccd1?w=400&h=400&fit=crop",  # pencil skirt
        "LS006": "https://images.unsplash.com/photo-1485462537746-965f33f7f6a7?w=400&h=400&fit=crop",  # midi skirt
    }
    
    print(f"Fixing {len(fixes)} placeholder images...")
    
    # Get current data
    df = table.search().limit(100).to_pandas()
    
    # Update the specific items
    for article_id, new_url in fixes.items():
        print(f"\nFixing {article_id}...")
        
        # Find the item in the dataframe
        item_mask = df['article_id'] == article_id
        if not item_mask.any():
            print(f"Item {article_id} not found!")
            continue
            
        # Download and encode the new image
        new_image_data = download_and_encode_image(new_url)
        if new_image_data:
            # Update the image_data for this item
            df.loc[item_mask, 'image_data'] = new_image_data
            print(f"✅ Updated {article_id} with new image")
        else:
            print(f"❌ Failed to update {article_id}")
        
        # Small delay to be respectful
        time.sleep(0.5)
    
    # Drop and recreate the table with updated data
    try:
        db.drop_table(table_name)
        print(f"Dropped existing table: {table_name}")
    except Exception as e:
        print(f"Table {table_name} couldn't be dropped: {e}")
    
    # Create new table with updated data
    print(f"Recreating table {table_name} with fixed images...")
    table = db.create_table(table_name, df)
    
    print(f"✅ Successfully updated table '{table_name}' with fixed placeholder images!")
    print("\n🖼️ All placeholder images have been replaced with real images!")
    print(f"📡 You can test the fixes at: http://localhost:8000/image/<article_id>")
    print(f"🌐 Examples:")
    for article_id in fixes.keys():
        print(f"  - http://localhost:8000/image/{article_id}")

if __name__ == "__main__":
    fix_placeholder_images()
