# 简介
一些大模型相关的工具链.

## Gradio
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

## Webviz
Webviz 是一个基于 Web 的开源可视化工具, 主要由 Cruise 公司开发, 用于回放、分析和可视化机器人操作系统 ROS 的 Bag 数据文件. 允许用户无需安装复杂软件, 只需在浏览器中打开即可查看自动驾驶传感器数据(相机、激光雷达)和诊断信息, 旨在提高机器人数据分析和调试的效率.
- 浏览器运行: 无需安装 ROS 环境, 打开 .bag 文件即可可视化
- ROS 兼容性好: 适合 ROS Bag 文件数据分析, 提供丰富的可视化组件、3D 场景、二维绘图、相机图
- 实时监控: 利用 ROS Bridge 对正在运行的机器人进行实时监控
- 团队协作: 基于 Web 工具便于团队成员远程访问和共享可视化
- 开源与定制: 可自定义布局

## 数据集格式
### 机器人 VLA 数据集格式
机器人领域的数据非常复杂, 包含图像、关节角度、末端姿势、夹爪状态、语言指令等. 常见的数据集格式有:
- RLDS / TFRecord: 将数据视为序列决策, 以 Episode 为单位, 每个 Episode 包含多个 Step. 每个 Step 包含 observation、action、reward、discount 等. 读写依赖 tensorflow 或者特定的库. RLDS 是用于存储和处理强化学习/机器人轨迹数据的标准和生态系统, RLDS 一般运行在 TFRecord 之上, TFRecord 是一种二进制文件格式.
- HDF5: 开源社区最常用的格式, 它采用分层数据结构, 通常包含 action 数组、observations / images 组、observations / qpos(关节位置) 等.
- Dexdata: 视频存为 .mp4、文本数据存为 .jsonl, 通过 index_cache.json 索引.

### 自动驾驶 VLA 数据集格式
自动驾驶 VLA 更关注车辆的控制(速度、转向)和轨迹预测, 而非机械臂的关节控制.
- QA 格式: 问答对齐格式, 在端到端自动驾驶 VLA (如 SANA 标准)中, 数据通常被处理成问答对的形式. 比如输入图像、历史轨迹, 输出未来轨迹预测、驾驶指令(左转、减速等).
- Nuscenes / Waymo 格式: 主要是感知数据集, VLA 模型通常基于这些格式进行微调. 其核心是多传感器融合(激光雷达、摄像头、雷达), 配合高精地图和标注框. 在这个基础上, 增加自然语言指令和轨迹规划.

### 通用多模态/基础模型格式
训练通用"大脑"的基础大模型, 数据格式更接近 NLP 或 CV 的标准.
- Alpaca: 适用于指令微调, 结构特点是 instruction、input、output, 让模型学会听懂人话.
- ShareGPT: 适用于多轮对话, 包含 conversations 列表, 包含 human 和 gpt 的交替对话.
- LMDB + JSONL: 适用于大模型训练, 将图片打包成 tar 包, 元数据(Caption 或 QA 对) 存为 jsonl.

## Onyx 
onyx 是一款开源可自托管的开源企业级 AI 聊天平台, 专门用于构建基于内部数据的知识库和搜索系统. 它有以下核心特性:
- 企业级知识库连接器(40+ Data Sources): 支持超过 50 数据源, 包括 slack、github、notion 等分散的文件.
- 深度学习与智能检索(Agentic RAG): 与传统的向量数据库不一样, onyx 使用基于代理的检索增强生成, 根据用户问题自主决定使用的语义检索、关键词匹配或知识图谱, 确保答案的准确性.
- 私有化部署与高隐私: 支持私有化部署, 数据安全性高.
- 全模型支持: 包括 OpenAI、Anthropic 等云端模型, Ollama、LM Studio 等本地模型.

应用场景如下:
- 企业内网知识助手: 员工可以通过对话方式在数万份文档(PDF、DOCX、Markdown) 中定位信息, 不需要记忆标题.
- 技术文档与代码查询: 可以对接 github、confluence, 快速检索技术文档、代码仓库.
- 智能化研究助手: 利用 Deep Research 功能, 自动化进行多步搜索、信息集成和深度分析.

相比传统知识库的核心优势有如下:
- 接入难度低: 50+ 连接器, 直接接入各类知识.
- 搜索精度: 传统知识库主要采用单一向量检索, 容易出错. onyx 采用混合检索、知识图谱和 Agentic RAG.
- 隐私保护: 离线私有化部署, 数据安全.
- 多模态: 支持文档、图像、代码解释器的多模态数据

### 组件分析
Onyx Standard 模式的组件有如下:
- index 索引服务: 使用 vespaengine/vespa, 是 Onyx 系统的语义大脑和核心记忆库. 上传文档的时候, 会将文本切块, 通过模型服务器转化成向, 它会通过 ANN 技术快速找到意思相近但字面不同的内容. 它还支持混合检索, 将向量检索、关键词搜索结合起来. vespa 会将处理好的文本块、向量数据、倒排索引持久化存储到磁盘上.
- 代码解释与执行: 提供了 onyxdotapp/code-interprete, LLM 不擅长的数学计算、数据分析、图表绘制或逻辑推演任务, 会使用这个服务, 通过 Python 运行后获取结果.
- 搜索服务: 使用 opensearchproject/opensearch 提供搜索、分析. onyx 使用它快速检索关键词, 向量不仅存储在 vespa, 也会将原始文件、标题、元数据存入 OpenSearch. 提问的时候, OpenSearch 可以毫秒级匹配. OpenSearch 提供的还是传统高效的倒排索引搜索能力.
- 模型服务: onyxdotapp/onyx-model-server 提供模型索引、推理, 它会加载嵌入模型或大语言模型. 文档基于这个组件的向量模型如 nomic-embed-text 进行向量化交给 vespa 服务, 由 indexing_model_server 提供. inference_model_server 服务则负责加载大语言模型 LLM 进行实时对话, 将 vespa 检索的文档片段结合问题生成回答.
- Web 交互: onyxdotapp/onyx-web-server 提供类 ChatGPT 的交互界面, 主要是 Next.js + React + TypeScript 编写的前端应用. 它调用后台 onyx-backend 的 API, 通过 WebSocket 长连接完成 Chat 实时对话. onyx-backend 提供 api_server 和 background 两个服务, 其中 api_server 提供 restful api, 实时处理前端界面的请求. background 核心任务是在后台默默处理耗时多的工作, 不阻塞前端交互, 比如定时处理 onyx 连接的外部系统 slack、github 等, 将文档(PDF、Word 等)进行文本提取、切片、调用 embedding 模型生成向量, 存入向量数据库建立索引等耗时多的工作.