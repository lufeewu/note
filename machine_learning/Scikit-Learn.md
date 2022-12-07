# 简介
sklearn 是开源库

## 库函数
sklearn 提供了许多可以直接使用的类和库函数。

**StandardScaler 类**: 数据标准化(也称为数据归一化)类, 通过类似 z=(x-u)/s 可以将 x 标准化为 z. StandardScaler 中提供了 fit、fit_transform、get_feature_names_out、get_params、inverse_transform、partial_fit、set_params、transform 等 method.
- fit(X, y=None, sample_weight=None): 计算将用于缩放的平均数和标准差.
- transform(X, copy=None): 通过集中和缩放执行标准化.

## SVM 模型


## 参考资料
1. [scikit-learn StandardScaler](https://scikit-learn.org/stable/modules/generated/sklearn.preprocessing.StandardScaler.html)
