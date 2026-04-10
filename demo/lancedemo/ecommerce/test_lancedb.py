#!/usr/bin/env python3
"""
Test script for LanceDB integration
"""

import sys
import os

# Add current directory to path for imports
sys.path.append(os.path.dirname(os.path.abspath(__file__)))

def test_lancedb_connection():
    """Test basic LanceDB connection"""
    try:
        import lancedb
        print("✓ LanceDB imported successfully")
        
        # Test connection
        db = lancedb.connect("./data")
        print("✓ LanceDB connection successful")
        
        # Test table operations
        try:
            table = db.open_table("h&m-mini")
            print("✓ Table opened successfully")
        except:
            print("ℹ Table doesn't exist yet (this is normal for first run)")
        
        return True
        
    except ImportError as e:
        print(f"✗ Failed to import LanceDB: {e}")
        return False
    except Exception as e:
        print(f"✗ LanceDB connection failed: {e}")
        return False

def test_sentence_transformers():
    """Test sentence transformers"""
    try:
        from sentence_transformers import SentenceTransformer
        print("✓ Sentence Transformers imported successfully")
        
        # Test encoder
        encoder = SentenceTransformer("clip-ViT-B-32")
        test_text = "test fashion item"
        embedding = encoder.encode(test_text)
        print(f"✓ Embedding generated successfully (dimension: {len(embedding)})")
        
        return True
        
    except ImportError as e:
        print(f"✗ Failed to import Sentence Transformers: {e}")
        return False
    except Exception as e:
        print(f"✗ Sentence Transformers test failed: {e}")
        return False

def test_fastapi():
    """Test FastAPI"""
    try:
        from fastapi import FastAPI
        print("✓ FastAPI imported successfully")
        return True
    except ImportError as e:
        print(f"✗ Failed to import FastAPI: {e}")
        return False

def main():
    """Run all tests"""
    print("Testing LanceDB Integration")
    print("=" * 30)
    
    tests = [
        ("LanceDB", test_lancedb_connection),
        ("Sentence Transformers", test_sentence_transformers),
        ("FastAPI", test_fastapi)
    ]
    
    passed = 0
    total = len(tests)
    
    for test_name, test_func in tests:
        print(f"\nTesting {test_name}...")
        if test_func():
            passed += 1
        else:
            print(f"✗ {test_name} test failed")
    
    print(f"\n{'=' * 30}")
    print(f"Tests passed: {passed}/{total}")
    
    if passed == total:
        print("🎉 All tests passed! You're ready to run the application.")
        print("\nTo start the application, run:")
        print("  ./start.sh          # On macOS/Linux")
        print("  start.bat           # On Windows")
        print("  # Or manually:")
        print("  uvicorn app:app --reload --host 0.0.0.0 --port 8000")
    else:
        print("❌ Some tests failed. Please check the errors above.")
        print("\nMake sure you have installed all dependencies:")
        print("  pip install -r requirements.txt")
    
    return passed == total

if __name__ == "__main__":
    success = main()
    sys.exit(0 if success else 1)
