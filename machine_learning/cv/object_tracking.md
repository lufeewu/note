# 简介
深度学习, 目标追踪相关.

## 目标追踪
通过 dlib 提供的算法库, 可以训练目标追踪模型. 然后从视频流中对目标进行追踪. 已经训练好的 demo 可以参考 [github - dlib-object-tracking](https://github.com/LaggyHammer/dlib-object-tracking).

## Caffe 框架
[caffe](https://caffe.berkeleyvision.org/) 是一个深度学习框架, 它具有易于上手、速度快、模块化的特性.

### 使用
在 Mac M2 上可以通过如下命令在 docker 运行: 
    
    docker run --platform linux/amd64 -ti bvlc/caffe:cpu caffe --version

可以参考 [Training LeNet on MNIST with Caffe](https://caffe.berkeleyvision.org/gathered/examples/mnist.html) 训练 LeNet 模型.

### web demo
提供 bvlc_reference_caffenet 模型, 该模型是由 AlexNet 模型改写. 为图片进行分类. 支持 1000+ 类目.
- ImageNet: 是一个计算机视觉数据集, 由斯坦福大学的李飞飞教授创建. 含 1000w+ 图片 2w+ Synet 索引.
- ILSVRC: ImageNet Large Scale Visual Recognition Challenge (ILSVRC) 使用 ImageNet 的一个子集, 基于 ImageNet 数据集的 1000 个类别的比赛. 

### python 库
提供了 python 的 caffe 库, 可以用于训练模型, 使用模型推理.
- caffe.Classifier 用于分类 

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
+ caffeModel: caffe 模型文件, 可以用来进行图像分类、目标检测、语音识别、自然语言处理、模型剪枝与压缩等. [caffemodel](http://dl.caffe.berkeleyvision.org/) 可以下载开源模型. caffe 提供了许多训练好的 caffe_model, 在 [Model Zoo](https://github.com/BVLC/caffe/wiki/Model-zoo)里管理.

### 架构与原理
Caffe 框架有 5 个组件, Blob、Solver、Net、Layer、Proto.
- Solver: 它协调整个模型的运行, 主要作用是交替调用前向(forward)和后向(backward)算法来更新参数, 从而最小化 loss, 是一种迭代的优化算法. 
- Blob: Blob 是 caffe 中的基础数据结构, 主要作用是存储和传输数据, 隐藏 CPU 和 GPU 之间的数据同步的细节.
- Net: Net 是由一些列层组成的有向无环图 DAG, 一个典型的 Net 从 data layer 开始, 从 loss layer 结束.
- Layer: layer 层主要定义三种运算, setup、forward 和 backward, 在 layer 内部主要是正向传导(forward)和反向传导(backward).

## 参考
1. [计算机视觉项目: 用dlib进行单目标跟踪](https://www.atyun.com/31701.html)
2. [github - dlib ](https://github.com/davisking/dlib)
3. [github - dlib-object-tracking](https://github.com/LaggyHammer/dlib-object-tracking)
4. [Caffemodel: 深度学习领域的经典模型](https://developer.baidu.com/article/details/1848415)
5. [Caffe: Things to know to train your network](https://github.com/arundasan91/Deep-Learning-with-Caffe/blob/master/Caffe_Things_to_know.md)
6. [机器学习周志华 pdf](https://github.com/Mikoto10032/DeepLearning/blob/master/books/机器学习周志华.pdf)
7. [aws - 什么是深度学习](https://aws.amazon.com/cn/what-is/deep-learning)
8. [Caffe - Layers](https://caffe.berkeleyvision.org/tutorial/layers.html)
9. [Caffe - docker](https://github.com/BVLC/caffe/tree/master/docker)
10. [MNIST数据集](https://docs.ultralytics.com/zh/datasets/classify/mnist/)
11. [LeNet](https://paddlepedia.readthedocs.io/en/latest/tutorials/computer_vision/classification/LeNet.html)
12. [Gradient-based learn- ing applied to document recognition](http://yann.lecun.com/exdb/publis/pdf/lecun-01a.pdf)
13. [caffe - Web Demo](https://caffe.berkeleyvision.org/gathered/examples/web_demo.html)
14. [ImageNet Classification with Deep Convolutional Neural Networks](https://papers.nips.cc/paper_files/paper/2012/hash/c399862d3b9d6b76c8436e924a68c45b-Abstract.html)
15. [caffemodel](http://dl.caffe.berkeleyvision.org/)
