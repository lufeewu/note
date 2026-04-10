#!/bin/bash

echo "H&M Fashion Search with LanceDB"
echo "================================"

# Check if virtual environment exists
if [ ! -d "venv" ]; then
    echo "Creating virtual environment..."
    python3 -m venv venv
fi

# Activate virtual environment
echo "Activating virtual environment..."
source venv/bin/activate

# Install dependencies
echo "Installing dependencies..."
pip install -r requirements.txt

# Create data directory if it doesn't exist
if [ ! -d "data" ]; then
    echo "Creating data directory..."
    mkdir -p data
fi

# Check if sample data exists
if [ ! -f "data/.sample_data_loaded" ]; then
    echo "Loading sample data..."
    python load_sample_data.py
    touch data/.sample_data_loaded
fi

# Start the application
echo "Starting the application..."
echo "Open http://localhost:8000/static/index.html in your browser"
echo "Press Ctrl+C to stop the server"
echo ""

uvicorn app:app --reload --host 0.0.0.0 --port 8000
