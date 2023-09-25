# 简介
大数据相关.

## 知识
- DAG 图: 有向无环图(Directed Acyclic Graph) 的简称. 在大数据处理中, DAG 计算常指将计算任务在内部分解成若干个子任务, 将子任务之间的逻辑关系或顺序构建成 DAG(有向无环图)结构.
- 图数据库: 使用图数据进行存储, 同时使用图结构进行语义查询的数据库. 它能高效地将关联数据的实体作为顶点(vertex)存储, 关系作为边(edge)存储, 并允许对这些边结构进行高性能的检索和查询, 也可以为这些点和边添加属性. 图数据库十分适合处理复杂关系的数据.
    + 高性能: 相较于关系型数据库和其它非关系型数据库, 在处理深度关联数据时, 性能更佳.
    + 灵活: 提供了极其灵活的数据模型, 可以根据业务变化实时对数据模型进行修改, 数据库的设计者无需计划数据库未来用例的详细信息.
    + 敏捷: 图数据库的数据建模非常直观, 支持测试驱动开发模式, 每次构建可进行功能测试和性能测试, 符合敏捷开发需求, 极大的提升生产和交付效率.

## Superset
Apache Superset 是 Airbnb 开源数据可视化软件. 主要提供了 Dashboard 和多维分析两大类功能. Superset 支持的数据源包括 CSV、MySQL、Oracle、Redshift、Drill、Hive、Impala、Elasticsearch 等 27 种数据源, 并深度支持 Druid.

类似的开源 BI 软件有:
- Grafana: 由 go 语言开发的开源数据可视化工具, 主要针对监控与日志分析, 带有告警功能.
- Metabase: 开源 BI 分析工具, 针对业务人员探索数据, 兼容大数据和传统数据库的分析工具.
- Redash: 是一款开源 BI 工具, 支持快速数据查询与可视化, 具备报警、订阅等功能.


## 数据集
[data.world](https://data.world/datasets/economics) 是一家数据分析和协作平台供应商, 上面有不同专业的人分享不同类型的数据.


## 资料
1. [深入浅出的实践大数据 DAG 图](https://xie.infoq.cn/article/4d4ab8c6a14577dd8c3ba465d)
2. [bi开源工具排行](https://juejin.cn/s/bi开源工具排行)
3. [data.world](https://data.world/datasets/economics)