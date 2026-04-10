import os
import lancedb
from lancedb.server import serve

# 1. 配置 MinIO 环境变量 (告诉 LanceDB 数据存在哪里)
# 注意：这里使用的是 s3:// 协议
os.environ["LANCEDB_URI"] = "s3://lance/"
os.environ["AWS_ACCESS_KEY_ID"] = "minioadmin"
os.environ["AWS_SECRET_ACCESS_KEY"] = "minioadmin123"
os.environ["AWS_ENDPOINT_URL"] = "http://minio:9000" # 如果是 Linux，可能需要换成 MinIO 的 IP
os.environ["AWS_ALLOW_HTTP"] = "true"
os.environ["AWS_S3_FORCE_PATH_STYLE"] = "true"
os.environ["AWS_REGION"] = "us-east-1"

# 2. 启动服务
# host="0.0.0.0" 允许外部访问
# port=8080 是 LanceDB 的默认端口
if __name__ == "__main__":
    print("🚀 Starting LanceDB Server on port 8080...")
    print("📦 Data will be stored in MinIO at s3://lancedb-bucket/")
    serve(host="0.0.0.0", port=8080)