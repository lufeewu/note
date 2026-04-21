# 简介
大预言模型 chatgpt 等相关.

## AI 知识
- 7b: 通常是 GPT-3 模型中的一个版本, 意思是 7 billion parameters, 是拥有 70 亿个参数的模型.

## ChatGPT
聊天生成预训练转换器(Chat Generative Pre-trained Transformer), 是 OpenAI 开发的人工智能聊天机器人. 程序基于 GPT-3.5 架构的大型语言模型并以强化学习训练.

- ChatGLM-6B: 开源、支持中英双语的对话语言模型. [THUDM/ChatGLM-6B](https://github.com/THUDM/ChatGLM-6B)
- Belle: BE Large Language model Engine 开源中文对话大模型. [LianjiaTechBelle](https://github.com/LianjiaTech/BELLE)
- LoRA: Low-Rank Adaptation of Large Language Models. [microsoft/LoRA](https://github.com/microsoft/LoRA)

## LLM
大型语言模型(LLM, Large Language Model)是基于大数据进行预训练的超大型深度学习模型. 底层转换器是一组神经网络, 这些神经网络由具有自注意力功能的编码器和解码器组成.
- 多模态大模型: 在一个统一的框架下, 集成了多种不同类型数据处理能力的大型神经网络模型. 这些模型能够处理图像、文本、音频等不同的数据模态, 并在这些模态之间进行有效的交互和信息整合. 多模态是指利用多种不同形式或感知渠道的信息进行表达、交流和理解的方式, 通常包括视觉、听觉、文本、触觉等多种感官输入和输出方式。在计算科学、人工智能和机器学习领域, 多模态技术指的是通过整合来自不同模态的数据(如图像、文字、音频、视频等)，从而增强模型的理解能力和推理能力。
- AGI(Artificial General Intelligence, AGI) 是具备与人类同等智能、或超越普通人类的人工智能, 能表现正常人类所具有的所有智能行为.

## LangChain
LangChain 是一个用于开发语言模型驱动的应用程序的框架。它使得应用程序能够:
- 具有上下文感知能力: 将语言模型连接到上下文源(提示指令, 少量的示例)
- 具有推理能力: 依赖语言模型进行推理(根据提供的上下文如何回答, 采取什么行动)

## MCP
MCP(Model Context Protocol, 模型上下文协议), 由 Anthropic 在 2024 年 11 月推出的一种开放标准, 旨在统一大语言模型(LLM)与外部数据源和工具之间的通信协议。主要目的是解决当前 AI 模型因为数据孤岛限制而无法充分发挥潜力的难题，MCP 使得 AI 应用能够安全地访问和操作本地及远程数据，为 AI 应用提供了万物链接的接口。

## MoE
MoE(全称 Mixture of Experts) 混合专家模型, 其核心工作设计思路是"术业有专攻", 将任务分门别类, 然后给多个"专家"进行解决。与 MoE 相对应的是稠密(Dense)模型,可以理解为一个通才模型。

## Agent 智能体
Agent 不仅具备语言理解能力和生成能力, 还可以利用工具, 如搜索引擎、数据库查询工具、邮件发送工具等与外界互动, 获取信息、执行操作。通过编排层根据目标自主规划行动, 合理地调用各种工具.一个完整的 Agent 通常由三个核心组件构成, 它们相互协作共同支撑 Agent 的智能行为:
- 模型(Model): ChatGpt 等，是 Agent 的大脑, 具有语言理解、推理、规划等能力.
- 工具(Tools): Agent 与外界交互的桥梁, 允许 Agent 访问外部数据库和服务、执行各种任务.
- 编排层(Orchestration Layer): Agent 的指挥中心，负责管理 Agent 的内部状态, 协调模型和工具的使用, 并根据目标指导 Agent 的行动.

## 参考
1. [wikipedia - ChatGPT](https://zh.wikipedia.org/zh-tw/ChatGPT)
2. [什么是大型语言模型？](https://aws.amazon.com/cn/what-is/large-language-model/)
3. [ollama 使用自己的微调模型](https://blog.csdn.net/spiderwower/article/details/138755776)
4. [一文看懂：MCP(大模型上下文协议)](https://zhuanlan.zhihu.com/p/27327515233)
5. [从零到手搓一个Agent: AI Agents新手入门精通（一）](https://cloud.tencent.com/developer/article/2487274)
6. [Awesome MCP Servers](https://github.com/punkpeye/awesome-mcp-servers)
7. [百度百科 - 多模态](https://baike.baidu.com/item/多模态/10647898)