# LangChain 对话机器人 Demo

基于 LangChain 实现的多轮对话机器人，支持流式输出和会话历史管理。

## 特性

- 多轮对话，自动维护对话历史
- 流式输出（Streaming）
- 支持 OpenAI 兼容接口（Ollama、通义千问、DeepSeek 等）
- 会话管理（新建/清空/查看历史）

## 快速开始

```bash
# 1. 安装依赖
pip install -r requirements.txt

# 2. 配置 API Key
cp .env.example .env
# 编辑 .env，填写你的 API Key

# 3. 运行
python chatbot.py
```

## 配置说明（.env）

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `OPENAI_API_KEY` | API Key | — |
| `OPENAI_BASE_URL` | API 地址，可替换为本地 Ollama 等 | `https://api.openai.com/v1` |
| `MODEL_NAME` | 模型名称 | `gpt-4o-mini` |
| `TEMPERATURE` | 温度（0~1） | `0.7` |

### 使用 Ollama（本地模型）

```env
OPENAI_API_KEY=ollama
OPENAI_BASE_URL=http://localhost:11434/v1
MODEL_NAME=qwen2.5:7b
```

## 指令

| 指令 | 说明 |
|------|------|
| `/help` | 显示帮助 |
| `/history` | 查看历史记录 |
| `/clear` | 清空当前会话 |
| `/new` | 开启新会话 |
| `/quit` | 退出 |
