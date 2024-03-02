# 简介
机器学习、深度学习相关知识.

## 概念
Tensor: 张量，深度学习的基本元素，类似高维度向量. 
derivative: 模型导数
gradient: 梯度
optimizer: 优化器，如 sgd、adam、rmsprop

Neural Network: 神经网络模型.
训练集: 用于训练神经网络模型的输入数据，MMNIST
error: 目标和预测结果的差距
epoch: 迭代完所有的训练数据为一个 epoch
loss function: 损失函数是将随机事件或有关随机变量的取值映射为非负实数用于表示该随机事件的"风险"或"损失"的函数. 常见的损失函数有 0-1 损失函数、绝对值损失函数、log对数损失函数、平方损失函数、指数损失函数、感知损失函数、交叉墒损失函数.

## 模型评估
通过模型评估, 直观的评价训练模型的性能.
### 混淆矩阵
混淆矩阵 TP、TN、FP、FN 可以用来评估分类模型.
- True Positive(真正, TP): 将正类预测为正类数.
- True Negative(真负, TN): 将负类预测为负类数.
- False Positive(假正, FP): 将负类预测为正类数 → 误报(Type I error).
- False Negative(假负 , FN): 将正类预测为负类数 → 漏报(Type II error).
- 精确率: P = TP/(TP+FP)
- 准确率: ACC = (TP+TN)/(TP+TN+FP+FN)
- 召回率: R = TP/(TP+FN)
- ROC 曲线: 在逻辑回归里, 设定一个阈值, 大于这个值为正类, 小于这个值为负类. 减小这个阈值, 那么更多的样本会被识别为正类, 可以提高正类识别率, 同时会有更多的负类被错误的识别为正类. 可以用 ROC 曲线评价分类器的好坏. 通过 TPR(纵坐标) = TP/(TP+FN) 和 FPR(横坐标) = FP/(FP+TN) 两个指标描绘曲线.
- AUC(Area Under Curve) : AUC 被定义为 ROC 曲线下的面积, AUC 值越大的分类器, 正确率越高. 0.5 < AUC < 1 的情况下, 分类器优于随机猜测, 有预测价值.

## Bert
Bidirectional Encoder Representation from Transformers 是 2018 年 Google AI 研究院提出的预训练模型, 使用的是多层 Transformer 结构. 

## Transformer
Transformer 是一个利用注意力机制来提高模型训练速度的模型, 主要用于自然语言处理(NLP) 与计算机视觉(CV) 领域. Transformer 模型旨在处理自然语言等顺序输入数据, 可用于翻译、文本摘要等任务.

## 因果推断
因果推断是一门研究因果关系的跨学科领域, 其核心内容是确定给定条件下变量之间的因果关系, 以及在特定干预下如何解释和估计这种关系.
+ 双重差分法: Differences-in-Differences(DID) 别名倍差法, 常用于政策效应评估.
+ Double Machine Learning: Double Machine Learning(DML) 是一种基于观测数据进行因果建模去偏方法.
+ 相关性生成机制: 因果、混淆偏差、样本选择偏差
+ 因果关系阶梯: 关联和预测、干预、反事实推理. 其中干预常用的方法是 AB test, 主要维度有 ATE、CATE、ITE .
+ 因果推断流派: 潜在因果模型、因果图模型、计量经济学方法, 其中因果模型和因果图模型是互联网数据分析的主流的两个流派.

## prompt 
prompt 是指在使用机器学习模型时为输入添加的一段文本或指令. 目的是为了引导模型生成更准、更有针对性的输出. 它可以是一个问题、一段描述、一种格式化的输入.

## 神经网络
神经网络是一种模仿生物神经网络的结构和功能的数学模型或计算模型.
### 神经元
神经元是构成神经网络的基本元素.
### 感知机
感知机是由两层神经元构成的.

### LMDB
LMDB(Lightning Memory-Mapped Database), 闪存映射嵌入式数据库. 是内存效率极高的数据库. 具有纯内存数据库的读取性能, 也保留了磁盘数据库的持久性. 训练模型数据集 MNIST 从 LMDB 数据库里读取图像数据 data 和标签数据 label.

### Datum
Datum 提供了创建、读取 tfrecord 数据集作为 tf.data.Datasets 的 api 集合.

## 题目
1. LR 推导、求导、梯度更新
2. SVM 原形式、对偶形式
3. FM 公式推导
4. GBDT 手推
5. XGB 推导
6. AUC 计算
7. 神经网络的反向传播
8. pytorch 编写 DNN
9. 常见的评价指标(准确率、混淆矩阵、精确率、召回率、F1 值、ROC-AUC、P-R 曲线、MAE、MSE、RMSE、R-square、AUC、MAP、NDCG、IDGC、MRR、轮廓系数、兰德指数、互信息)

## 资料
1. <a src="https://github.com/wnzhang/deep-ctr">Deep Learning for Ad CTR Estimation</a>
2. <a src="https://github.com/shenweichen/DeepCTR/">DeepCTR</a>
3. <a src="https://github.com/INTERMT/Awesome-PyTorch-Chinese">PyTorch 中文学习资料集合</a>
4. <a src="http://charleshm.github.io/2016/03/Model-Performance/">机器学习性能评估指标</a>
5. [梳理常见机器学习面试题](https://zhuanlan.zhihu.com/p/82105066)
6. [BERT](https://paddlepedia.readthedocs.io/en/latest/tutorials/pretrain_model/bert.html)
7. [知乎 - 十分钟理解Transformer](https://zhuanlan.zhihu.com/p/82312421)
8. [知乎 - 双重差分法（DID）介绍](https://zhuanlan.zhihu.com/p/48952513)
9. [维基百科 - 人工神经网络](https://zh.wikipedia.org/wiki/人工神经网络)