#!/usr/bin/env python3
"""
Sample data loader for H&M Fashion Search with LanceDB
This script creates sample fashion data and loads it into the LanceDB table.
"""

import pandas as pd
import lancedb
from transformers import CLIPModel, CLIPTokenizer
import numpy as np
import torch
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

def create_sample_data():
    """Create sample fashion data"""
    sample_data = [
        # Original 5 items
        {
            "image_url": "https://placehold.co/300x400/4A90E2/FFFFFF?text=Shirt",
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
            "image_url": "https://placehold.co/300x400/E24A90/FFFFFF?text=Dress",
            "prod_name": "Elegant Summer Dress",
            "detail_desc": "A beautiful summer dress with floral pattern, perfect for warm weather. Made from lightweight, breathable fabric.",
            "product_type_name": "Dress",
            "index_group_name": "Ladieswear",
            "price": 49.99,
            "article_id": "LD001",
            "available": True,
            "color": "Floral",
            "size": "S"
        },
        {
            "image_url": "https://placehold.co/300x400/2E8B57/FFFFFF?text=Jacket",
            "prod_name": "Athletic Performance Jacket",
            "detail_desc": "High-performance athletic jacket designed for outdoor sports. Features moisture-wicking technology and adjustable fit.",
            "product_type_name": "Jacket",
            "index_group_name": "Sport",
            "price": 79.99,
            "article_id": "SJ001",
            "available": True,
            "color": "Black",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/FF6B6B/FFFFFF?text=Sweater",
            "prod_name": "Cozy Kids Sweater",
            "detail_desc": "Warm and comfortable sweater for children, perfect for cold weather. Made from soft, hypoallergenic materials.",
            "product_type_name": "Sweater",
            "index_group_name": "Baby/Children",
            "price": 24.99,
            "article_id": "BC001",
            "available": True,
            "color": "Red",
            "size": "4T"
        },
        {
            "image_url": "https://placehold.co/300x400/666666/FFFFFF?text=Gloves",
            "prod_name": "Fashion Winter Gloves",
            "detail_desc": "Stylish winter gloves with touchscreen compatibility. Features warm lining and modern design for the fashion-conscious.",
            "product_type_name": "Gloves",
            "index_group_name": "Divided",
            "price": 19.99,
            "article_id": "DG001",
            "available": True,
            "color": "Gray",
            "size": "M"
        },
        
        # Additional 45 items - Menswear (15 items)
        {
            "image_url": "https://placehold.co/300x400/1E90FF/FFFFFF?text=T-Shirt",
            "prod_name": "Premium Cotton T-Shirt",
            "detail_desc": "Soft premium cotton t-shirt with modern fit. Perfect for casual wear and layering. Available in multiple colors.",
            "product_type_name": "T-Shirt",
            "index_group_name": "Menswear",
            "price": 19.99,
            "article_id": "MS002",
            "available": True,
            "color": "White",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/000080/FFFFFF?text=Jeans",
            "prod_name": "Slim Fit Denim Jeans",
            "detail_desc": "Classic slim fit jeans made from premium denim. Features comfortable stretch and modern styling for everyday wear.",
            "product_type_name": "Jeans",
            "index_group_name": "Menswear",
            "price": 59.99,
            "article_id": "MS003",
            "available": True,
            "color": "Dark Blue",
            "size": "32"
        },
        {
            "image_url": "https://placehold.co/300x400/8B4513/FFFFFF?text=Chinos",
            "prod_name": "Smart Casual Chinos",
            "detail_desc": "Versatile chino pants perfect for smart casual occasions. Made from cotton twill with a comfortable fit.",
            "product_type_name": "Chinos",
            "index_group_name": "Menswear",
            "price": 39.99,
            "article_id": "MS004",
            "available": True,
            "color": "Khaki",
            "size": "34"
        },
        {
            "image_url": "https://placehold.co/300x400/708090/FFFFFF?text=Polo",
            "prod_name": "Classic Polo Shirt",
            "detail_desc": "Timeless polo shirt made from breathable pique cotton. Features classic collar and two-button placket.",
            "product_type_name": "Polo",
            "index_group_name": "Menswear",
            "price": 34.99,
            "article_id": "MS005",
            "available": True,
            "color": "Gray",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/2F4F4F/FFFFFF?text=Blazer",
            "prod_name": "Business Casual Blazer",
            "detail_desc": "Sophisticated blazer perfect for business casual settings. Tailored fit with premium fabric blend.",
            "product_type_name": "Blazer",
            "index_group_name": "Menswear",
            "price": 129.99,
            "article_id": "MS006",
            "available": True,
            "color": "Charcoal",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/DC143C/FFFFFF?text=Hoodie",
            "prod_name": "Comfortable Pullover Hoodie",
            "detail_desc": "Cozy pullover hoodie with kangaroo pocket. Made from soft cotton blend for maximum comfort.",
            "product_type_name": "Hoodie",
            "index_group_name": "Menswear",
            "price": 44.99,
            "article_id": "MS007",
            "available": True,
            "color": "Red",
            "size": "XL"
        },
        {
            "image_url": "https://placehold.co/300x400/228B22/FFFFFF?text=Shorts",
            "prod_name": "Summer Casual Shorts",
            "detail_desc": "Lightweight summer shorts perfect for warm weather. Features elastic waistband and side pockets.",
            "product_type_name": "Shorts",
            "index_group_name": "Menswear",
            "price": 24.99,
            "article_id": "MS008",
            "available": True,
            "color": "Olive",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/4169E1/FFFFFF?text=Vest",
            "prod_name": "Formal Dress Vest",
            "detail_desc": "Elegant dress vest for formal occasions. Features adjustable back strap and premium fabric construction.",
            "product_type_name": "Vest",
            "index_group_name": "Menswear",
            "price": 49.99,
            "article_id": "MS009",
            "available": True,
            "color": "Royal Blue",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/800080/FFFFFF?text=Cardigan",
            "prod_name": "Wool Blend Cardigan",
            "detail_desc": "Classic cardigan made from premium wool blend. Features button closure and ribbed details.",
            "product_type_name": "Cardigan",
            "index_group_name": "Menswear",
            "price": 69.99,
            "article_id": "MS010",
            "available": True,
            "color": "Purple",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/FF4500/FFFFFF?text=Tank",
            "prod_name": "Athletic Tank Top",
            "detail_desc": "Breathable tank top perfect for workouts and hot weather. Moisture-wicking fabric with comfortable fit.",
            "product_type_name": "Tank Top",
            "index_group_name": "Menswear",
            "price": 16.99,
            "article_id": "MS011",
            "available": True,
            "color": "Orange",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/FFD700/000000?text=Henley",
            "prod_name": "Long Sleeve Henley",
            "detail_desc": "Classic henley shirt with button placket. Made from soft cotton for everyday comfort and style.",
            "product_type_name": "Henley",
            "index_group_name": "Menswear",
            "price": 32.99,
            "article_id": "MS012",
            "available": True,
            "color": "Gold",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/20B2AA/FFFFFF?text=Flannel",
            "prod_name": "Flannel Button-Up Shirt",
            "detail_desc": "Cozy flannel shirt perfect for cooler weather. Features classic plaid pattern and soft brushed finish.",
            "product_type_name": "Flannel",
            "index_group_name": "Menswear",
            "price": 42.99,
            "article_id": "MS013",
            "available": True,
            "color": "Teal",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/B22222/FFFFFF?text=Sweater",
            "prod_name": "Crew Neck Sweater",
            "detail_desc": "Classic crew neck sweater made from merino wool. Perfect for layering or wearing alone in cool weather.",
            "product_type_name": "Sweater",
            "index_group_name": "Menswear",
            "price": 89.99,
            "article_id": "MS014",
            "available": True,
            "color": "Burgundy",
            "size": "XL"
        },
        {
            "image_url": "https://placehold.co/300x400/8FBC8F/FFFFFF?text=Tracksuit",
            "prod_name": "Athletic Tracksuit Set",
            "detail_desc": "Complete tracksuit set with jacket and pants. Made from moisture-wicking fabric for active lifestyle.",
            "product_type_name": "Tracksuit",
            "index_group_name": "Menswear",
            "price": 79.99,
            "article_id": "MS015",
            "available": True,
            "color": "Light Green",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/F0E68C/000000?text=Tie",
            "prod_name": "Silk Dress Tie",
            "detail_desc": "Premium silk tie perfect for formal occasions. Features classic pattern and luxurious feel.",
            "product_type_name": "Tie",
            "index_group_name": "Menswear",
            "price": 24.99,
            "article_id": "MS016",
            "available": True,
            "color": "Gold",
            "size": "One Size"
        },
        
        # Ladieswear (15 items)
        {
            "image_url": "https://placehold.co/300x400/FF69B4/FFFFFF?text=Blouse",
            "prod_name": "Silk Chiffon Blouse",
            "detail_desc": "Elegant silk chiffon blouse with delicate draping. Perfect for office wear or special occasions.",
            "product_type_name": "Blouse",
            "index_group_name": "Ladieswear",
            "price": 59.99,
            "article_id": "LD002",
            "available": True,
            "color": "Pink",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/9370DB/FFFFFF?text=Skirt",
            "prod_name": "A-Line Midi Skirt",
            "detail_desc": "Classic A-line midi skirt with flattering fit. Made from premium fabric with comfortable waistband.",
            "product_type_name": "Skirt",
            "index_group_name": "Ladieswear",
            "price": 34.99,
            "article_id": "LD003",
            "available": True,
            "color": "Purple",
            "size": "S"
        },
        {
            "image_url": "https://placehold.co/300x400/00CED1/FFFFFF?text=Jumpsuit",
            "prod_name": "Wide Leg Jumpsuit",
            "detail_desc": "Stylish wide leg jumpsuit with belt detail. Perfect for both casual and dressy occasions.",
            "product_type_name": "Jumpsuit",
            "index_group_name": "Ladieswear",
            "price": 79.99,
            "article_id": "LD004",
            "available": True,
            "color": "Turquoise",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/FF1493/FFFFFF?text=Top",
            "prod_name": "Cropped Tank Top",
            "detail_desc": "Trendy cropped tank top perfect for summer. Made from soft cotton with comfortable fit.",
            "product_type_name": "Top",
            "index_group_name": "Ladieswear",
            "price": 22.99,
            "article_id": "LD005",
            "available": True,
            "color": "Hot Pink",
            "size": "S"
        },
        {
            "image_url": "https://placehold.co/300x400/32CD32/FFFFFF?text=Cardigan",
            "prod_name": "Lightweight Cardigan",
            "detail_desc": "Versatile lightweight cardigan perfect for layering. Features open front design and soft knit fabric.",
            "product_type_name": "Cardigan",
            "index_group_name": "Ladieswear",
            "price": 44.99,
            "article_id": "LD006",
            "available": True,
            "color": "Lime Green",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/8A2BE2/FFFFFF?text=Pants",
            "prod_name": "High Waist Trousers",
            "detail_desc": "Elegant high waist trousers with wide leg cut. Perfect for office wear or evening occasions.",
            "product_type_name": "Pants",
            "index_group_name": "Ladieswear",
            "price": 54.99,
            "article_id": "LD007",
            "available": True,
            "color": "Blue Violet",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/FF6347/FFFFFF?text=Sweater",
            "prod_name": "Cozy Knit Sweater",
            "detail_desc": "Soft and cozy knit sweater perfect for cool weather. Features relaxed fit and ribbed details.",
            "product_type_name": "Sweater",
            "index_group_name": "Ladieswear",
            "price": 49.99,
            "article_id": "LD008",
            "available": True,
            "color": "Coral",
            "size": "S"
        },
        {
            "image_url": "https://placehold.co/300x400/4682B4/FFFFFF?text=Blazer",
            "prod_name": "Professional Blazer",
            "detail_desc": "Tailored professional blazer with structured shoulders. Perfect for business meetings and formal events.",
            "product_type_name": "Blazer",
            "index_group_name": "Ladieswear",
            "price": 89.99,
            "article_id": "LD009",
            "available": True,
            "color": "Steel Blue",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/DA70D6/FFFFFF?text=Romper",
            "prod_name": "Summer Romper",
            "detail_desc": "Cute summer romper with adjustable straps. Perfect for beach days and casual outings.",
            "product_type_name": "Romper",
            "index_group_name": "Ladieswear",
            "price": 39.99,
            "article_id": "LD010",
            "available": True,
            "color": "Orchid",
            "size": "S"
        },
        {
            "image_url": "https://placehold.co/300x400/CD5C5C/FFFFFF?text=Tunic",
            "prod_name": "Flowy Tunic Top",
            "detail_desc": "Comfortable flowy tunic with three-quarter sleeves. Perfect for casual wear and layering.",
            "product_type_name": "Tunic",
            "index_group_name": "Ladieswear",
            "price": 36.99,
            "article_id": "LD011",
            "available": True,
            "color": "Indian Red",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/40E0D0/FFFFFF?text=Camisole",
            "prod_name": "Silk Camisole",
            "detail_desc": "Luxurious silk camisole with delicate lace trim. Perfect for layering or wearing alone.",
            "product_type_name": "Camisole",
            "index_group_name": "Ladieswear",
            "price": 42.99,
            "article_id": "LD012",
            "available": True,
            "color": "Turquoise",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/F4A460/FFFFFF?text=Wrap+Dress",
            "prod_name": "Wrap Style Dress",
            "detail_desc": "Flattering wrap style dress with tie waist. Made from jersey fabric for comfort and style.",
            "product_type_name": "Wrap Dress",
            "index_group_name": "Ladieswear",
            "price": 64.99,
            "article_id": "LD013",
            "available": True,
            "color": "Sandy Brown",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/DDA0DD/FFFFFF?text=Leggings",
            "prod_name": "High Waist Leggings",
            "detail_desc": "Comfortable high waist leggings with compression fit. Perfect for workouts or casual wear.",
            "product_type_name": "Leggings",
            "index_group_name": "Ladieswear",
            "price": 28.99,
            "article_id": "LD014",
            "available": True,
            "color": "Plum",
            "size": "S"
        },
        {
            "image_url": "https://placehold.co/300x400/98FB98/000000?text=Kimono",
            "prod_name": "Floral Kimono",
            "detail_desc": "Beautiful floral kimono perfect for layering. Features flowing silhouette and delicate print.",
            "product_type_name": "Kimono",
            "index_group_name": "Ladieswear",
            "price": 52.99,
            "article_id": "LD015",
            "available": True,
            "color": "Pale Green",
            "size": "One Size"
        },
        {
            "image_url": "https://placehold.co/300x400/F0F8FF/000000?text=Shawl",
            "prod_name": "Cashmere Shawl",
            "detail_desc": "Luxurious cashmere shawl perfect for evening wear. Soft and elegant with beautiful drape.",
            "product_type_name": "Shawl",
            "index_group_name": "Ladieswear",
            "price": 89.99,
            "article_id": "LD016",
            "available": True,
            "color": "Alice Blue",
            "size": "One Size"
        },
        
        # Sport (8 items)
        {
            "image_url": "https://placehold.co/300x400/00FF00/000000?text=Sports+Bra",
            "prod_name": "High Support Sports Bra",
            "detail_desc": "High support sports bra for intense workouts. Features moisture-wicking fabric and comfortable straps.",
            "product_type_name": "Sports Bra",
            "index_group_name": "Sport",
            "price": 34.99,
            "article_id": "SP002",
            "available": True,
            "color": "Lime",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/FF8C00/FFFFFF?text=Joggers",
            "prod_name": "Athletic Joggers",
            "detail_desc": "Comfortable athletic joggers with tapered fit. Features elastic waistband and side pockets.",
            "product_type_name": "Joggers",
            "index_group_name": "Sport",
            "price": 44.99,
            "article_id": "SP003",
            "available": True,
            "color": "Orange",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/1E90FF/FFFFFF?text=Yoga+Pants",
            "prod_name": "High Waist Yoga Pants",
            "detail_desc": "Stretchy yoga pants with high waistband. Perfect for yoga, pilates, and general fitness activities.",
            "product_type_name": "Yoga Pants",
            "index_group_name": "Sport",
            "price": 39.99,
            "article_id": "SP004",
            "available": True,
            "color": "Dodger Blue",
            "size": "S"
        },
        {
            "image_url": "https://placehold.co/300x400/DC143C/FFFFFF?text=Running+Shirt",
            "prod_name": "Moisture Wicking Running Shirt",
            "detail_desc": "Lightweight running shirt with moisture-wicking technology. Features reflective details for safety.",
            "product_type_name": "Running Shirt",
            "index_group_name": "Sport",
            "price": 29.99,
            "article_id": "SP005",
            "available": True,
            "color": "Crimson",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/FFD700/000000?text=Swimsuit",
            "prod_name": "One Piece Swimsuit",
            "detail_desc": "Stylish one piece swimsuit with chlorine-resistant fabric. Perfect for swimming and water sports.",
            "product_type_name": "Swimsuit",
            "index_group_name": "Sport",
            "price": 54.99,
            "article_id": "SP006",
            "available": True,
            "color": "Gold",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/800080/FFFFFF?text=Workout+Tank",
            "prod_name": "Mesh Panel Workout Tank",
            "detail_desc": "Breathable workout tank with mesh panels. Features loose fit for maximum comfort during exercise.",
            "product_type_name": "Workout Tank",
            "index_group_name": "Sport",
            "price": 26.99,
            "article_id": "SP007",
            "available": True,
            "color": "Purple",
            "size": "L"
        },
        {
            "image_url": "https://placehold.co/300x400/228B22/FFFFFF?text=Cycling+Shorts",
            "prod_name": "Padded Cycling Shorts",
            "detail_desc": "Professional cycling shorts with padded insert. Features compression fit and moisture management.",
            "product_type_name": "Cycling Shorts",
            "index_group_name": "Sport",
            "price": 49.99,
            "article_id": "SP008",
            "available": True,
            "color": "Forest Green",
            "size": "M"
        },
        {
            "image_url": "https://placehold.co/300x400/4169E1/FFFFFF?text=Windbreaker",
            "prod_name": "Lightweight Windbreaker",
            "detail_desc": "Packable windbreaker perfect for outdoor activities. Features water-resistant coating and ventilation.",
            "product_type_name": "Windbreaker",
            "index_group_name": "Sport",
            "price": 59.99,
            "article_id": "SP009",
            "available": True,
            "color": "Royal Blue",
            "size": "L"
        },
        
        # Baby/Children (4 items)
        {
            "image_url": "https://placehold.co/300x400/FFB6C1/000000?text=Onesie",
            "prod_name": "Cotton Baby Onesie",
            "detail_desc": "Soft cotton onesie with snap closure. Perfect for everyday wear and easy diaper changes.",
            "product_type_name": "Onesie",
            "index_group_name": "Baby/Children",
            "price": 12.99,
            "article_id": "BC002",
            "available": True,
            "color": "Light Pink",
            "size": "6M"
        },
        {
            "image_url": "https://placehold.co/300x400/87CEEB/FFFFFF?text=Kids+Pajamas",
            "prod_name": "Fun Print Pajama Set",
            "detail_desc": "Comfortable pajama set with fun animal prints. Made from soft cotton for peaceful sleep.",
            "product_type_name": "Pajamas",
            "index_group_name": "Baby/Children",
            "price": 19.99,
            "article_id": "BC003",
            "available": True,
            "color": "Sky Blue",
            "size": "5T"
        },
        {
            "image_url": "https://placehold.co/300x400/98FB98/000000?text=Kids+T-Shirt",
            "prod_name": "Graphic Kids T-Shirt",
            "detail_desc": "Fun graphic t-shirt for active kids. Features colorful design and comfortable cotton fabric.",
            "product_type_name": "T-Shirt",
            "index_group_name": "Baby/Children",
            "price": 14.99,
            "article_id": "BC004",
            "available": True,
            "color": "Pale Green",
            "size": "8"
        },
        {
            "image_url": "https://placehold.co/300x400/DDA0DD/FFFFFF?text=Kids+Dress",
            "prod_name": "Twirl Dress for Girls",
            "detail_desc": "Adorable twirl dress perfect for special occasions. Features full skirt and comfortable bodice.",
            "product_type_name": "Dress",
            "index_group_name": "Baby/Children",
            "price": 29.99,
            "article_id": "BC005",
            "available": True,
            "color": "Plum",
            "size": "6"
        },
        
        # Divided (3 items)
        {
            "image_url": "https://placehold.co/300x400/FF4500/FFFFFF?text=Beanie",
            "prod_name": "Trendy Knit Beanie",
            "detail_desc": "Stylish knit beanie perfect for cold weather. Features soft acrylic yarn and comfortable fit.",
            "product_type_name": "Beanie",
            "index_group_name": "Divided",
            "price": 14.99,
            "article_id": "DV002",
            "available": True,
            "color": "Orange Red",
            "size": "One Size"
        },
        {
            "image_url": "https://placehold.co/300x400/20B2AA/FFFFFF?text=Scarf",
            "prod_name": "Infinity Scarf",
            "detail_desc": "Cozy infinity scarf in soft knit fabric. Perfect for adding warmth and style to any outfit.",
            "product_type_name": "Scarf",
            "index_group_name": "Divided",
            "price": 22.99,
            "article_id": "DV003",
            "available": True,
            "color": "Light Sea Green",
            "size": "One Size"
        },
        {
            "image_url": "https://placehold.co/300x400/8B4513/FFFFFF?text=Belt",
            "prod_name": "Leather Belt",
            "detail_desc": "Classic leather belt with metal buckle. Perfect for completing any casual or formal outfit.",
            "product_type_name": "Belt",
            "index_group_name": "Divided",
            "price": 29.99,
            "article_id": "DV004",
            "available": True,
            "color": "Brown",
            "size": "M"
        }
    ]
    
    data_df = pd.DataFrame(sample_data)
    # Keep schema compatible with app.py image endpoint.
    data_df["image_data"] = ""
    return data_df


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

def generate_embeddings(data_df, text_column='detail_desc'):
    """Generate embeddings for text descriptions"""
    print("Generating embeddings...")
    encoder = CLIPTextEncoder(Config.EMBEDDING_MODEL)
    embeddings = encoder.encode_batch(data_df[text_column].fillna(""))
    embeddings = [row.tolist() for row in embeddings]

    data_df['vector'] = embeddings
    return data_df

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


def load_data_to_lancedb(data_df, db_path=None, table_name=Config.TABLE_NAME):
    """Load data into LanceDB table"""
    if db_path is None:
        db_path = _build_lancedb_path(Config.USE_MINIO)

    print(f"Connecting to LanceDB at {db_path}...")
    db = lancedb.connect(db_path)
    
    try:
        # Try to open existing table
        table = db.open_table(table_name)
        print(f"Opened existing table: {table_name}")
    except Exception as e:
        message = str(e).lower()
        table_missing = (
            "not found" in message
            or "does not exist" in message
            or "no such table" in message
            or "table" in message and "missing" in message
        )
        if table_missing:
            # Create new table only when it's truly missing.
            print(f"Creating new table: {table_name}")
            table = db.create_table(table_name, data_df)
            print("Table created successfully!")
            return
        raise RuntimeError(f"Failed to open table '{table_name}' on MinIO: {e}")
    
    # Add data to existing table
    print("Adding data to existing table...")
    table.add(data_df)
    print("Data added successfully!")

def main():
    """Main function to load sample data"""
    print("H&M Fashion Search - Sample Data Loader")
    print("=" * 50)
    
    # Create sample data
    print("Creating sample fashion data...")
    data_df = create_sample_data()
    
    # Generate embeddings
    data_df = generate_embeddings(data_df)
    
    # Load into LanceDB
    load_data_to_lancedb(data_df)
    
    print("\nSample data loaded successfully!")
    print(f"Loaded {len(data_df)} fashion items")
    if Config.USE_MINIO:
        print(f"LanceDB path: s3://{Config.MINIO_BUCKET}/{Config.MINIO_PATH}")
    print("\nYou can now start the application with:")
    print("uvicorn app:app --reload --host 0.0.0.0 --port 8000")

if __name__ == "__main__":
    main()
