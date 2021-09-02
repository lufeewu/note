# 简介
etcd 是一个数据存储在磁盘的 kv 存储器，它实现了 raft 算法从而达到分布式一致性。

## etcd 能力
etcd 提供了 key-value 存储的能力、监听机制、key 过期、续约、原子 CAS、CAD 等。

### 存储能力
etcd 提供的是 key-value 的存储能力，使用 B+ 树进行存储。B+ 树会存储于内存，以提高查询效率。

### Watch 机制
etcd 提供了 Watch 机制，用于监听一组或者一个 key 的键值变动。

### CAS 与 CAD
CAS 与 CAD 主要用于确保修改、删除的原子性。
- CAS(Compare and Swap), 通过比较 old value 与实际的 value 值，确认无变化后再替换成 new value，确保原子性。
- CAD(Compare and Delete), 当 old value 和实际的 value 相等时，删除 key，不相等则不删除。

### 性能
etcd 的官方宣称性能是提供  10000 writes/sec 的写性能。

## raft 算法
etcd 在分布式一致性上使用了 raft 算法。raft 算法是一种 multi-paxos，它强化了 leader 的地位。

### 选举
raft 算法通过中心化的 leader 确保分布式一致性。而在高可用上，通过 leader 选举机制确保半数节点正常工作时，集群整体是可以正常提供服务的。

### 脑裂
raft 通过 region leader 算法保证一致性。



## 参考
1. [CAS和CAD命令- 云数据库Redis 版 - 阿里云帮助文档](https://help.aliyun.com/apsara/enterprise/v_3_14_0_20210519/kvstore/enterprise-product-introduction/cas-and-cad-commands.html)