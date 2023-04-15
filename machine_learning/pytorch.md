## 简介
PyTorch 是一个开源 Python 机器学习库, 基于 Torch , 底层由 c++ 实现. 是学术界实现深度学习算法最常用的框架.

## 概念
主要模块如下:
- Torch 张量: 类似 Numpy 数组的多维矩阵, 可用于高效计算.
- torch.autograd 自动求导: 自动计算模型参数的导数, 用于反向传播更新权值.
- 神经网络 nn 模块: 实现常用的神经网络层, 可以方便地搭建模型.
- torch.optim 模块: 提供各种优化算法, 用于训练模型.
- torch.utils 模块: 提供数据加载和预处理工具等辅助功能. 
- torch.cuda: 是 PyTorch 的 GPU 加速模块, 能方便的在 GPU 上运行.

## 面试题
1. MSELoss、CTCLoss、BCELoss 函数有什么用?
2. 反向传播是什么?
3. Pytorch 的基本要素是什么?

## 参考
1. [pytorch框架图](https://juejin.cn/s/pytorch框架图)