# 简介
ClickHouse 是一个列式数据库管理系统(DMBS), 数据在 ClickHouse 中始终是列式存储的.

## 概念
- **列式存储**: Online Analytical Processing, OLAP 将数据的每一列组织在一起, 大大减少在进行聚合计算时候磁盘 I/O 次数, 但在写入时要多次 I/O.
- **行模式存储**: Online Transaction Processing, OLTP 数据基于行存储, 数据的写入快, 但对于需要大量聚合统计数据的需求效率并不高.
- **列**: 表示内存中的列, 需使用 IColumn 接口.
- **字段**: 表示
- **抽象漏洞**: 
- **数据类型**: IDataType 负责序列化和番序列化.
- **块**: Block 表示内存中的表的子集的容器.
- **块流**: 快流用于处理数据.
- **格式**: 有用于向客户端输出数据的格式等, 如 Pretty、TabSeparated、JsonEachRow 等, 还有行流 IRowInputStream 等.
- **I/O**: 用于面向字节的输入输出, ReadBuffer、WriteBuffer 抽象类.
- **表**: 又 IStorage 接口表示, 接口的不同实现对应不同的表引擎, 如 StorageMergeTree、StorageMemory 等.
- **解析器**: 
- **解释器**: 
- **函数**: 
- **聚合函数**: 
- **服务器**: 
- **合并树**: 
- **复制**: 

## 参考
1. [什么是列式存储，一文秒懂](https://juejin.cn/post/6844904118872440840)
2. [github - ClickHouse/ClickHouse](https://github.com/ClickHouse/ClickHouse)