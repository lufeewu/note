# 简介
ClickHouse 是一个开源列式数据库管理系统(DMBS), 数据在 ClickHouse 中始终是列式存储的. 它用于实时数据分析, 处理速度比传统方法快 100-1000 倍.

## 概念
- **列式存储**: Online Analytical Processing, OLAP 将数据的每一列组织在一起, 大大减少在进行聚合计算时候磁盘 I/O 次数, 但在写入时要多次 I/O.
- **行模式存储**: Online Transaction Processing, OLTP 数据基于行存储, 数据的写入快, 但对于需要大量聚合统计数据的需求效率并不高.
- **列**: 表示内存中的列, 需使用 IColumn 接口.
- **字段**: 用于表示单个值.
- **抽象漏洞**: Leaky Abstractions, IColumn 拥有常见的用于关系型数据转换的方法.
- **数据类型**: IDataType 负责序列化和番序列化.
- **块**: Block 表示内存中的表的子集的容器.
- **块流**: 快流用于处理数据.
- **格式**: 有用于向客户端输出数据的格式等, 如 Pretty、TabSeparated、JsonEachRow 等, 还有行流 IRowInputStream 等.
- **I/O**: 用于面向字节的输入输出, ReadBuffer、WriteBuffer 抽象类.
- **表**: 又 IStorage 接口表示, 接口的不同实现对应不同的表引擎, 如 StorageMergeTree、StorageMemory 等.
- **解析器**: Parsers 是一个手写递归下降解析器, 用于解析查询
- **解释器**: 负责创建 AST 查询执行流水线.
- **函数**: 有普通函数和聚合函数. 函数作用在以 block 为单位的数据上, 以实现向量查询执行.
- **聚合函数**: 是状态函数, 它们将传入的值激活到某个状态, 并允许从该状态获取结果. 使用 IAggregateFunction 进行管理.
- **服务器**: Server 实现了多个不同的接口. 用于外部客户端的 HTTP 接口、用于本机 Clickhouse 客户端及分布式查询的 TCP 接口、用于传输数据进行拷贝的接口.
- **合并树**: MergeTree 是一系列支持按主键索引的存储引擎.
- **复制**: 基于表实现的复制, 可以在同一个服务器上有可复制的表和不可复制的表.

### 架构
ClickHouse 的主要组件包括 Client、Server、Storage、Distributed、ZooKeeper、Replication .
- **Client**: 用于与 ClickHouse 服务器进行通信, 向其发送查询请求.
- **Server**: 服务器, 是 ClickHouse 的核心组件, 负载接受来自客户端的查询请求, 将结果返回给客户端.
- **Storage**: 存储, ClickHouse 支持多种存储引擎, 包括 MergeTree、Collapsing 等.
- **Distributed**: 支持分布式部署, 可以用多个节点组成一个 ClickHouse 集群, 提高系统的可用性和性能.
- **Zookeeper**: 是一个分布式的协调服务, ClickHouse 可以通过 Zookeeper 实现集群节点的发现和管理.
- **Replication**: ClickHouse 支持数据复制, 可以将数据从一个节点复制到其他节点, 提高系统的可用性和数据的容错性.

## 性能
相比于 MySQL、InfluxDB, ClickHouse 在同样的资源情况导入速度、磁盘占用、查询性能方面都十分突出. 
- 4c16g 的资源, 官网 6600w 数据集, 导入耗时约 75 秒、磁盘空间 2.7G、全表 count 100 ms、 max/min 186ms、平均值 123ms、方差 113ms.
- clickhouse 快的原因: 列式存储、数据压缩、向量化执行引擎、多线程和分布式、多样的表引擎(MergeTree).
- clickhouse 缺点: 不支持事务、不支持真删除/更新、分布式能力弱、不支持高并发(建议 QPS 100).

## 参考
1. [什么是列式存储，一文秒懂](https://juejin.cn/post/6844904118872440840)
2. [github - ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse)
3. [ClickHouse 架构概述](https://clickhouse.com/docs/zh/development/architecture)
4. [记一次 ClickHouse 性能测试](https://juejin.cn/post/7131778389865660452)
5. [ClickHouse教程](https://clickhouse.com/docs/zh/getting-started/tutorial)