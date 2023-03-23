# 简介
Apache Spark 是开源集群运算框架.允许用户将资料加载至集群存储器内，并进行多次查询，十分适合机器学习算法.它主要使用内存中的缓存和优化的查询执行方式，可针对任何规模的数据进行快速分析查询，提供 java、python 等语言 API.支持跨多个工作负载重用代码，包括批处理、交互式查询、实时分析、机器学习和图形处理等.

## 基础概念
工作负载:
- 平台基础 Spark Core: 负责内存管理、故障恢复、计划安排、分配与监控作业，以及和存储系统进行交互.
- 交互查询 SparkSQL: 低延迟交互式查询的分布式查询引擎，速度可比 MapReduce 快 100 倍.
- 实时分析 Spark Streaming: 利用 Spark Core 的快速计划功能流式分析的实时解决方案.
- 机器学习 Spark MLib: MLib 是在大规模数据上进行机器学习所需的算法库.算法包括分类、回归、集群、协同过滤和模式挖掘等功能.
- 图形处理 Spark GraphX: GraphX 是构建在 Spark 上的分布式图形处理框架.提供 ELT、探索性分析和迭代图形计算，让用户能够以交互方式大规模构建、转换图形数据结构.

- Spark RDD: 是一种数据存储集合.只能由它支持的数据源或是其它 RDD 经过一定的转换产生.在 RDD 上可以执行的操作有两种转换和行动，每个 RDD 都记录了自己如何由持久化存储的源数据计算得出的，即血统.
- HDFS: 适合存储大文件存储的分布式文件系统. 
- HBase: Hbase 是 Hadoop 数据库, 一个分布式、可伸缩的大数据存储. Hbase 是在 HDFS 的基础之上构建的.
    + 列式存储(RowKey): 充当了主键的作用, 可以唯一的标识一行记录.
<img src="../img/hbase_struct.webp">

### yarn
yarn 是一个资源调度平台，负责为运算程序提供服务器运算资源，相当于分布式的操作系统平台.
### PySpark
pyspark 提供 Spark 的 Python API，能够通过 python 操作 RDDs. Py4j 库允许 python 动态的操作 JVM 的对象.

### 表
- 基本表: 独立存在的表, 在 SQL 中一个关系对应一个表.
- 中间表: 用来兼容数据, 建立映射关系, 兼容新老数据表的数据. 一般是实体之间存在多对多的关系时, 创建一个中间表给实体建立联系. 中间表主要和 OLAP 业务有关, 主要是由于计算逻辑复杂、查询性能差、ETL 过程转存、多样性数据源混合计算等原因造成.
- 临时表: 
- 数据运营层(ODS): Operation Data Store 数据准备区, 也称为帖源层. 数据仓库源头系统的数据表通常会原封不动的存储一份, 称为 ODS 层, 是后续数据仓库加工数据的来源. 
- 数据仓库层(DW): DW 数据分层由下到上为 DWD、DWB、DWS.
    + 细节数据层(DWD): data warehouse details 细节数据层, 是业务层与数据仓库的隔离层. 主要对ODS数据层做一些数据清洗和规范化的操作.
    + 数据基础层(DWB): data warehouse base 数据基础层, 存储的是客观数据, 一般用作中间层, 可以认为是大量指标的数据层.
    + 数据服务层(DWS): data warehouse service 数据服务层, 基于DWB上的基础数据, 整合汇总成分析某一个主题域的服务数据层, 一般是宽表. 用于提供后续的业务查询, OLAP分析, 数据分发等.
- 数据服务层/应用层(ADS): Application DataService 应用数据服务, 该层主要是提供数据产品和数据分析使用的数据, 一般会存储在 ES, MySQL 等系统中供线上系统使用.

## 参考资料
1. [SparkSQL并行执行多个Job的探索](https://cloud.tencent.com/developer/article/1901879)
2. [AWS: 介绍 Apache Spark](https://aws.amazon.com/cn/big-data/what-is-spark/)
3. [Spark on Kubernetes 与 Spark on Yarn 不完全对比分析](https://www.infoq.cn/article/7cmvdianctkck4birhvi)
4. [Difference Between HDFS and HBase](https://www.educba.com/hdfs-vs-hbase/)
5. [中间表是什么？和报表有什么关系？会带来怎样的问题？又如何解决？](https://zhuanlan.zhihu.com/p/148782827)
6. [数据仓库分层中的ODS、DWD、DWS](https://www.cnblogs.com/amyzhu/p/13513425.html)