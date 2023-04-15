## 简介
tensorflow 是一个基于 python 的用于创建机器学习应用程序的库, 是执行复杂数学的低级工具包. 为用户提供了构建实验学习架构和可定制性选项.

## 概念
- TensorBoard: 一套可视化工具, 用于检查和理解 TensorFlow 运行和图表. 
- Client: 前端系统的主要组成部分, 支持多语言的编程环境, 基于计算图的编程模型, 方便用户构造复杂计算图, 实现各种形式的模型设计.
- Master: 在分布式的运行时环境中, 根据 Session.run 的 Fetching 参数, 从计算图中反向遍历, 找到所依赖的最小子图.
- Worker: 每个任务 TensorFlow 将启动一个 Worker Service, Worker Service 将按照计算图中节点之间的依赖关系, 根据当前可用的硬件环境(GPU/CPU), 调用 OP 的 Kernel 实现完成 OP 的运算.
- Kernel: 是 OP 在某种硬件环境的特定实现, 它负责执行 OP 的运算.
- 算法: tensorflow 支持多种算法. 包括回归分析、分类、卷积神经网络(CNN)、循环神经网络(RNN)、生成对抗网络(GANs)、卷积单元(Convolutional Units) 等, 还支持其它一些特殊算法, 如自编码(Autoencoder)、聚类(Clustering)等.

## kreas
kreas 是一个用 python 编写的开源神经网络库.

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