# 简介
Apache Spark 是开源集群运算框架。允许用户将资料加载至集群存储器内，并进行多次查询，十分适合机器学习算法。它主要使用内存中的缓存和优化的查询执行方式，可针对任何规模的数据进行快速分析查询，提供 java、python 等语言 API。支持跨多个工作负载重用代码，包括批处理、交互式查询、实时分析、机器学习和图形处理等。

## 基础概念
工作负载:
- 平台基础 Spark Core: 负责内存管理、故障恢复、计划安排、分配与监控作业，以及和存储系统进行交互。
- 交互查询 SparkSQL: 低延迟交互式查询的分布式查询引擎，速度可比 MapReduce 快 100 倍。
- 实时分析 Spark Streaming: 利用 Spark Core 的快速计划功能流式分析的实时解决方案。
- 机器学习 Spark MLib: MLib 是在大规模数据上进行机器学习所需的算法库。算法包括分类、回归、集群、协同过滤和模式挖掘等功能。
- 图形处理 Spark GraphX: GraphX 是构建在 Spark 上的分布式图形处理框架。提供 ELT、探索性分析和迭代图形计算，让用户能够以交互方式大规模构建、转换图形数据结构。

Spark RDD: 是一种数据存储集合。只能由它支持的数据源或是其它 RDD 经过一定的转换产生。在 RDD 上可以执行的操作有两种转换和行动，每个 RDD 都记录了自己如何由持久化存储的源数据计算得出的，即血统。
HDFS: 适合存储大文件存储的分布式文件系统. 
HBase: 
### yarn
yarn 是一个资源调度平台，负责为运算程序提供服务器运算资源，相当于分布式的操作系统平台。
### PySpark
pyspark 提供 Spark 的 Python API，能够通过 python 操作 RDDs. Py4j 库允许 python 动态的操作 JVM 的对象。

## 参考资料
1. [SparkSQL并行执行多个Job的探索](https://cloud.tencent.com/developer/article/1901879)
2. [AWS: 介绍 Apache Spark](https://aws.amazon.com/cn/big-data/what-is-spark/)
3. [Spark on Kubernetes 与 Spark on Yarn 不完全对比分析](https://www.infoq.cn/article/7cmvdianctkck4birhvi)
3. [Difference Between HDFS and HBase](https://www.educba.com/hdfs-vs-hbase/)