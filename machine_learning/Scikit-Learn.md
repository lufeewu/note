# 简介
sklearn 是开源库

## 库函数
sklearn 提供了许多可以直接使用的类和库函数。

**StandardScaler 类**: 数据标准化(也称为数据归一化)类, 通过类似 z=(x-u)/s 可以将 x 标准化为 z. StandardScaler 中提供了 fit、fit_transform、get_feature_names_out、get_params、inverse_transform、partial_fit、set_params、transform 等 method.
- fit(X, y=None, sample_weight=None): 计算将用于缩放的平均数和标准差.
- transform(X, copy=None): 通过集中和缩放执行标准化, 将样本数据归一化.


**SVC 类**: class sklearn.svm.SVC 是 C-Support Vector Classfication. 它基于 libsvm 实现的支持向量机分类器. 它的拟合时间与样本数量的二次方成正比. 创建 SVC 类的核心参数包括选取 svm 的核函数(包括 linear、poly、rbf、sigmoid、precomputed 或者自定义回调)、回归 C 参数等.

SVM 核函数:
- linear: 线性核函数, 在数据线性可分情况下使用，运算速度快，效果好, 但它不能处理线性不可分数据.
- poly: 多项式核函数, 可以将数据从低维空间映射到高维空间.
- rbf: 高斯核函数, 可以将样本映射到高维空间中.
- sigmoid: sigmoid 核函数经常用在神经网络的映射中，选用 sigmoid 核函数时候，svm 实现的是多层神经网络.

## SVM 模型
支持向量机(support vector machines, SVM) 是一种二分类模型，基本模型是定义在特征空间上的间隔最大的线性分类器.

## 参考资料
1. [scikit-learn StandardScaler](https://scikit-learn.org/stable/modules/generated/sklearn.preprocessing.StandardScaler.html)
