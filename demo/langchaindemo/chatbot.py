"""
LangChain 对话机器人 Demo
支持多轮对话，带历史记忆，兼容 OpenAI 兼容接口（如 Ollama、通义千问等）
"""

import os
from dotenv import load_dotenv
from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage, AIMessage, SystemMessage
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain_core.chat_history import BaseChatMessageHistory
from langchain_core.runnables.history import RunnableWithMessageHistory
from langchain_core.chat_history import InMemoryChatMessageHistory

load_dotenv()

# ── 配置 ─────────────────────────────────────────────────────────────────────

SYSTEM_PROMPT = """你是一个友好、专业的 AI 助手。
你的回答应该简洁、准确，并用中文回复（除非用户用其他语言提问）。
如果你不确定某个答案，请诚实地说明。"""

MODEL_NAME   = os.getenv("MODEL_NAME", "gpt-4o-mini")
BASE_URL     = os.getenv("OPENAI_BASE_URL", "https://api.openai.com/v1")
API_KEY      = os.getenv("OPENAI_API_KEY", "")
TEMPERATURE  = float(os.getenv("TEMPERATURE", "0.7"))

# ── 构建 LLM & Chain ──────────────────────────────────────────────────────────

def build_chain():
    llm = ChatOpenAI(
        model=MODEL_NAME,
        base_url=BASE_URL,
        api_key=API_KEY,
        temperature=TEMPERATURE,
        streaming=True,
    )

    prompt = ChatPromptTemplate.from_messages([
        SystemMessage(content=SYSTEM_PROMPT),
        MessagesPlaceholder(variable_name="history"),
        ("human", "{input}"),
    ])

    chain = prompt | llm
    return chain


# ── 会话历史管理 ──────────────────────────────────────────────────────────────

_session_store: dict[str, BaseChatMessageHistory] = {}

def get_session_history(session_id: str) -> BaseChatMessageHistory:
    if session_id not in _session_store:
        _session_store[session_id] = InMemoryChatMessageHistory()
    return _session_store[session_id]


# ── 对话函数 ──────────────────────────────────────────────────────────────────

def chat(chain_with_history, session_id: str, user_input: str) -> str:
    """发送消息并获取流式回复，返回完整回复文本。"""
    print("\nAssistant: ", end="", flush=True)
    full_response = ""
    for chunk in chain_with_history.stream(
        {"input": user_input},
        config={"configurable": {"session_id": session_id}},
    ):
        text = chunk.content
        print(text, end="", flush=True)
        full_response += text
    print()
    return full_response


def print_history(session_id: str):
    """打印当前会话历史。"""
    history = get_session_history(session_id)
    messages = history.messages
    if not messages:
        print("（暂无历史记录）")
        return
    for msg in messages:
        role = "You" if isinstance(msg, HumanMessage) else "AI"
        print(f"  [{role}]: {msg.content[:80]}{'...' if len(msg.content) > 80 else ''}")


def print_help():
    print("""
指令:
  /quit, /exit  — 退出
  /clear        — 清空当前会话历史
  /history      — 查看历史记录
  /new          — 开启新会话
  /help         — 显示帮助
""")


# ── 主入口 ────────────────────────────────────────────────────────────────────

def main():
    if not API_KEY or API_KEY == "your-api-key-here":
        print("[警告] 未设置 OPENAI_API_KEY，请在 .env 文件中配置。")
        print("  cp .env.example .env  # 然后填写你的 API Key\n")

    print("=" * 50)
    print("   LangChain 对话机器人 Demo")
    print(f"   模型: {MODEL_NAME}")
    print("   输入 /help 查看指令")
    print("=" * 50)

    chain = build_chain()
    chain_with_history = RunnableWithMessageHistory(
        chain,
        get_session_history,
        input_messages_key="input",
        history_messages_key="history",
    )

    session_id = "session-1"
    session_count = 1

    while True:
        try:
            user_input = input("\nYou: ").strip()
        except (EOFError, KeyboardInterrupt):
            print("\n再见！")
            break

        if not user_input:
            continue

        # 处理指令
        if user_input.startswith("/"):
            cmd = user_input.lower()
            if cmd in ("/quit", "/exit"):
                print("再见！")
                break
            elif cmd == "/clear":
                get_session_history(session_id).clear()
                print("已清空当前会话历史。")
            elif cmd == "/history":
                print_history(session_id)
            elif cmd == "/new":
                session_count += 1
                session_id = f"session-{session_count}"
                print(f"已开启新会话: {session_id}")
            elif cmd == "/help":
                print_help()
            else:
                print(f"未知指令: {user_input}，输入 /help 查看帮助。")
            continue

        # 正常对话
        try:
            chat(chain_with_history, session_id, user_input)
        except Exception as e:
            print(f"\n[错误] {e}")
            print("请检查 API Key 和网络连接是否正常。")


if __name__ == "__main__":
    main()
