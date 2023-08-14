# 简介
Apache Flink 是一个框架和分布式处理引擎, 用于在无边界和有边界数据流上进行有状态的计算.

## 概念
+ 处理有界和无界数据
    - 有界流: 有定义流的开始，也有定义流的结束. 可以在摄取所有数据后再进行计算. 有界流处理通常被称为批处理.
    - 无界流: 有定义流的开始，没有定义流的结束. 会无休止的产生数据, 无界流的数据需要持续处理.
+ 易部署: 集成了常见的资源管理器, 如 Hadoop YARN、Apache Mesos 或 kubernetes, 同时也可以作为独立集群允许.
+ 大规模: Flink 可以方便的扩展到大规模集群中. 用于维护万亿级事件、TB 大小的状态, 允许在千万级内核上, Flink 将任务并行化为数千个子任务, 充分利用集群的 CPU、内存、磁盘和网络资源.

## Flink 架构
Flink 主要针对实时计算领域处理流数据, 在此之前还有 Storm、SparkStreaming 等. 下面是几个实时计算框架的对比: 
![实时计算框架对比](../img/flink_compare.png)
- JobManager: 负责整个 Flink 集群任务的调度及资源管理, 从客户端获取提交的应用, 然后根据集群中 TaskManager 上 TaskSlot 的使用情况, 为提交的应用分配 TaskSlot 资源并命令 TaskManager 从客户端获取的应用.
- TaskManager: 相当于集群的的 Slave 节点, 负责具体的任务执行和对应任务在每个节点上的资源申请和管理.
- Client 客户端: 负责将任务提交到集群, 与 JobManager 构建 Akka 连接, 然后将任务提交到 JobManager, 通过和 JobManager 之间进行交互获取任务执行的状态.


## 数据仓库架构
海量数据分析的技术架构经历了多个演进历程，从小时级发展到亚秒级。
### Lambda 架构
Lambda 架构(Lambda Architecture) 是由 Twitter 工程师南森·马茨提出的大数据处理架构. Lambda 架构使开发人员能够构建大规模分布式数据处理系统。它具有很好的灵活性和可扩展性，对硬件故障和人为失误有很好的容错性。
![lambda 架构](../img/lambda_architecture.png)
Lambda 架构由三层系统组成:
- 批处理层: Batch Layer, 存储管理主要数据集(不可变的数据集)和预先处理计算好的视图.
- 速度处理层: Speed Layer, 实时处理新来的大数据.
- 服务层: Serving Layer, 在批处理层和速度处理层处理完的结果都输出存储在服务层中, 服务层通过返回预先计算的数据视图或从速度层处理构建好数据视图来响应查询.

### kappa 架构
Lambda 架构满足了实时的需求，但带来了许多开发和运维工作。随着 Flink 等流处理引擎的出现，流处理技术变成熟了，LinkedIn 的 Jay Kreps 提出了 Kappa 架构, 它简化了 Lambda 架构
![kappa 架构对比](../img/kappa_compare.png)

## 近线系统
通常将系统分位在线系统和离线系统，近线系统则介于二者之间，有时效性要求但为在线系统提供已经计算好的离线数据。近线系统的主要目的是实时、快捷的挖掘热点事件，并输出热点事件特征供上层应用使用。


## 参考资料
1. [【大数据实战】Docker中Flink集群搭建](https://www.cnblogs.com/isuning/p/16214378.html)
2. [Apache Flink 是什么？](https://flink.apache.org/zh/what-is-flink/flink-architecture/)
3. [Demo：基于 Flink SQL 构建流式应用](https://wuchong.me/blog/2020/02/25/demo-building-real-time-application-with-flink-sql/)
4. [Chapter 4. Platform Architecture](https://www.oreilly.com/library/view/open-source-data/9781492074281/ch04.html)
5. [实时数仓之 Kappa 架构与 Lambda 架构（建议收藏！）](https://zhuanlan.zhihu.com/p/584255261)
6. [大数据Flink进阶（四）：Flink应用场景以及其他实时计算框架对比原创](https://cloud.tencent.com/developer/article/2241665)