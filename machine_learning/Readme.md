# 简介
机器学习、深度学习


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

## 资料
1. <a src="https://github.com/wnzhang/deep-ctr">Deep Learning for Ad CTR Estimation</a>
2. <a src="https://github.com/shenweichen/DeepCTR/">DeepCTR</a>
3. <a src="https://github.com/INTERMT/Awesome-PyTorch-Chinese">PyTorch 中文学习资料集合</a>
4. <a src="http://charleshm.github.io/2016/03/Model-Performance/">机器学习性能评估指标</a>
