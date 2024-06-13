## 简介
tensorflow 是一个基于 python 的用于创建机器学习应用程序的库, 是执行复杂数学的低级工具包. 为用户提供了构建实验学习架构和可定制性选项.

## 概念
- TensorBoard: 一套可视化工具, 用于检查和理解 TensorFlow 运行和图表.
- Client: 前端系统的主要组成部分, 支持多语言的编程环境, 基于计算图的编程模型, 方便用户构造复杂计算图, 实现各种形式的模型设计.
- Master: 在分布式的运行时环境中, 根据 Session.run 的 Fetching 参数, 从计算图中反向遍历, 找到所依赖的最小子图.
- Worker: 每个任务 TensorFlow 将启动一个 Worker Service, Worker Service 将按照计算图中节点之间的依赖关系, 根据当前可用的硬件环境(GPU/CPU), 调用 OP 的 Kernel 实现完成 OP 的运算.
- Kernel: 是 OP 在某种硬件环境的特定实现, 它负责执行 OP 的运算.
- 算法: tensorflow 支持多种算法. 包括回归分析、分类、卷积神经网络(CNN)、循环神经网络(RNN)、生成对抗网络(GANs)、卷积单元(Convolutional Units)等, 还支持其它一些特殊算法, 如自编码(Autoencoder)、聚类(Clustering)等.

## 实践
mac 中通过 docker 启动 tensorflow 环境.

## keras
keras 是一个用 python 编写的开源神经网络库.

## Bert
Bidirectional Encoder Representation from Transformers 是 2018 年 Google AI 研究院提出的预训练模型, 使用的是多层 Transformer 结构. 

### Tokenizer
Tokenizer 标记器, 目标是将文本转换为模型可以处理的数据, 模型只能处理数字, 因此 Tokenizer 需要将文本输入转换为数字输入.

    import tokenize
    from transformers import BertTokenizer 

BertTokenizer 是基于 BasicTokenizer 和 WordPieceTokenizer 的分词器.

### keras.Model
kreas.Model 全功能模型类, 由层组成, 它可以训练、评估、加载、保存, 甚至在多台机器上进行训练.


## Tensorflow Hub
Tensorflow Hub 是一个包含经过训练的机器学习模型的代码库, 模型稍作调整便可以部署到任何设备上. 只需要几行代码即可重复经过训练的模型, 例如 BERT 和 Faster R-CNN.

## 面试题
1. 张量是什么?
2. 张量有多少种类型?
3. TensorFlow 主要特点是什么?
4. TensorFlow 的优势是什么?
5. TensorFlow 有什么局限性/缺点?
6. TensorFlow 架构的三个工作组件是什么?
7. 何时会在 TensorFlow 中发现模型的过拟合情况?

## 参考
1. [TensorFlow面试题和答案(2023年收集更新)](https://www.yiibai.com/interview/3000)
2. [TensorFlow 架构与设计-四大组件【转】](https://www.cnblogs.com/ningskyer/articles/6481898.html)
3. [Tensorflow Hub](https://www.tensorflow.org/hub?hl=zh-cn)
4. [如何使用Bert预训练模型进行文本分类？](https://fuxi.163.com/database/1052)
5. [tokenize — Tokenizer for Python source](https://docs.python.org/3/library/tokenize.html)
6. [BertTokenizer](https://huggingface.co/transformers/v3.0.2/model_doc/bert.html#berttokenizer)
7. [transformer 中 tokenizer 的那些事](https://www.cnblogs.com/carolsun/p/16903276.html)
8. [keras_模型](https://www.tensorflow.org/guide/intro_to_modules?hl=zh-cn)