# 简介
k8s 学习笔记

## k8s 单机部署
在 docker desktop 应用中, 内嵌了一个 kubernetes 集群, 可以快速开启. 用于个人学习掌握 kubernetes .

## Kubernetes CRD 开发
CRD 称为自定义资源定义, 用于定义用户定义的资源, 可以遵循 k8s 的控制器开发规范基于 client-go 进行调用并实现 Informer、ResourceEventHandler、Workqueue 等组件逻辑.
- Informer 组件: 是 client-go 的核心工具包, 是一个带有本地缓存和索引机制的、可以注册 EventHandler 的 client, 本地缓存称为 Store, 索引称为 Index. Informer 是为了减轻 apiserver 交互的压力抽象出来的 cache 层.
- ResourceEventHandler 组件: 处于用户的 controller 代码中, 回调函数对资源对象的变化进行处理.
- Workqueue 组件: 工作队列, 它的主要功能在于标记和去重, 它支持三种队列分别是 FIFO 队列 Interface、延迟队列 DelayingInterface、限速队列 RateLimitingInterface.

## 参考文献
1. [Mac下kubernetes安装及搭建部署单机服务 ](https://www.cnblogs.com/lucky-yqy/p/14362312.html)
2. [k8s crd demo](https://github.com/domac/crddemo)