# 简介
目标追踪相关.

## 目标追踪
通过 dlib 提供的算法库, 可以训练目标追踪模型. 然后从视频流中对目标进行追踪. 已经训练好的 demo 可以参考 [github - dlib-object-tracking](https://github.com/LaggyHammer/dlib-object-tracking).

## Caffe 框架
[caffe](https://caffe.berkeleyvision.org/) 是一个深度学习框架, 它具有易于上手、速度快、模块化的特性.

## OpenCV 
在 OpenCV 的 3.4 版本中支持了多种深度学习框架, 包括 Caffe、Tensorflow、torch/pytorch、darknet 等. OpenCV 给出了常见神经网络的 C++ 和 Python 接口.

### dnn 模块
在 OpenCV 3.3 之后加入了 DNN 深度学习模块, 借用 OpenCV 的 dnn 接口可以实现推理计算, 但不能实现模型训练.
+ dnn.readNetFromCaffe: 从 caffe 框架格式中读取网络模型
+ .prototxt 文件: 定义每层的结构信息
    - layer type: caffe 神经网络层 type 的分类有 Vision、Recurrent、Common、Normalization、Activation、Neuron、Utility、Loss 等, 包含 LeRU、ReLU、Convolution、Premute、Flatten、PriorBox、Concat、Reshape、Softmax、DetectionOutput 等类型的神经网络层.
    - bottom/top: 输入输出数据.
+ caffeModel: caffe 模型文件, 可以用来进行图像分类、目标检测、语音识别、自然语言处理、模型剪枝与压缩等

### 深度学习
深度学习是仿照人脑建模的神经网络. 深度学习神经网络(或人工神经网络)是由计算机内部协同工作的多层人工神经元组成的. 深度学习是机器学习的子集, 是为了提高传统机器学习技术的效率.
- 输入层: 人工神经网络有几个向其输入数据的节点. 这些节点构成了系统的输入层.
- 隐藏层: 输入层处理数据并将其传递到神经网络中的更远层.
- 输出层: 输出层由输出数据的节点组成. 输出 "是" 或 "否" 答案的深度学习模型在输出层中只有两个节点.

## 参考
1. [计算机视觉项目: 用dlib进行单目标跟踪](https://www.atyun.com/31701.html)
2. [github - dlib ](https://github.com/davisking/dlib)
3. [github - dlib-object-tracking](https://github.com/LaggyHammer/dlib-object-tracking)
4. [Caffemodel: 深度学习领域的经典模型](https://developer.baidu.com/article/details/1848415)
5. [Caffe: Things to know to train your network](https://github.com/arundasan91/Deep-Learning-with-Caffe/blob/master/Caffe_Things_to_know.md)
6. [机器学习周志华 pdf](https://github.com/Mikoto10032/DeepLearning/blob/master/books/机器学习周志华.pdf)
7. [aws - 什么是深度学习](https://aws.amazon.com/cn/what-is/deep-learning)
8. [Caffe - Layers](https://caffe.berkeleyvision.org/tutorial/layers.html)