# 简介
一些大模型相关的工具链.

## gradio
gradio 是一个开源的 python 库, 用于快速构建机器学习模型、API 或任意 Python 函数的交互式 Web 界面. 它旨在让开发者无需任何前端开发经验(无需写 HTML、CSS 或 JS), 即可通过几行代码实现模型的前端界面生成、测试、展示和部署.

特性:
- 极致简化, 几行代码即可输出图像上传、文本框和输出组件, 快速创建 web 应用
- 支持多模态数据, 文本、图像、音频、视频、视频流等
- 不仅限于 ML 模型、可分享生成访问链接、可标注 flag、支持拖拽设计

主要用途:
- 展示机器学习模型, 为科研人员、AI 开发者提供展示成果的快速方式
- 快速模型测试与调试
- 广泛用于 Hugging Face Spaces 上的模型托管
- 案例: Stable Diffusion WebUI、LLaMa Factory、GPT Academic(科研工作流) 均用 Gradio 构建前端

## Pillow
Pillow 是 Python 图像处理库, 用于打开、操作和保存多种图像格式. 广泛用于图像裁剪、旋转、滤镜、缩略图生成等基础操作. 

可以通过 tensorflow 库、torchvision.transforms 等库将 PIL.Image 转换成张量.

## MLX-VLM
是一个基于苹果 MLX(Machine Learning eXperience) 框架开发的多模态大模型工具包, 专门用于在 Mac 上高效运行、量化和微调视觉语言模型(Vision Language Models, VLMs).

可以通过如下命令通过 mlx 启动模型:
```shell
        # 指定 Hugging Face 模型 ID, 自动下载模型
        mlx_vlm --model mlx-community/Qwen2-VL-7B-Instruct-4bit

        # 指定本地模型目录启动
        mlx_vlm --model /Users/yourname/models/Qwen2-VL-7B-Instruct-4bit
```


## Open Webui
Open WebUI(Ollama WebUI)是一个开源、自托管且功能丰富的 AI 应用平台, 旨在为本地大语言模型(LLM)提供类似 ChatGPT 的用户体验. 设计上支持完离线运行, 能够轻松连接 Ollama 或 OpenAI 兼容的 API, 具备内置 RAG(检索增强生成)功能.