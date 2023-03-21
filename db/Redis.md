# 简介
Redis 是广泛使用的分布式缓存中间件. 可用于队列、分布式锁、Key-Value 数据库等. 

## 知识点
关于 Redis，需要掌握基本数据结构、跳表、Hash 冲突解决、持久化、主从、分布式、性能等知识.

### 持久化
redis 的持久化方式主要是 AOF、RDB .
AOF: 保存 binlog，将所有写入命令及参数写入文件保存在磁盘.
- AOF Rewrite 机制: 随着时间的推移，AOF 文件会膨胀的比较厉害，频繁的写入 AOF 文件将影响性能.系统通过 AOF Rewrite 机制定期重写 AOF 文件，将数据库的数据以协议方式保存到新的 AOF 文件中，重写的 AOF 小于原 AOF 文件，从而达到减小 AOF 文件目的.但是 AOF 也会占用较多的内存，写入 AOF buffer 过程中会阻塞工作线程.
- REWRITE: 主线程中重写 AOF，阻塞工作线程，生产环境中几乎不使用.
- BGREWRITE: 在后台子进程中重写 AOF，不阻塞工作线程.

RDB: 将 Redis 的当前数据以快照二进制的方式保存在磁盘，通过 RDB 的快照恢复数据更快.

### 主从
Redis 的高可用灾备方案主要是主从方式.
- 一主一从: 主要用于灾备、故障恢复、负载均衡等.将主节点 master 的数据单向的复制到从节点 slave 节点.当主节点故障时，由从节点提供服务，另外也可以进行读写分离，从节点提供读服务，提高并发量.
- 多主多从: redis 的多主多从指的是 redis 的集群.每个主节点有一个或多个从节点，而集群的主节点之间通过分片操作，每个master节点存储一部分数据.
- 分片: 将数据拆分到多个 Redis 实例的过程，主要解决单机 Redis 容量有限的问题.数据将按一定规则分配到多台机器.采用多主多从，没个分区一个主多个从.
- 哨兵: Redis 哨兵主要解决的是主从中出现宕机的情况，哨兵将监控 redis 的节点，当其中的一个主节点出现宕机故障时，将自动将该 master 节点的某个 slave 节点升级为 master 节点，从而达到 redis 的高可用.

### 分片机制
Redis 通过一致性哈希解决分布式集群中，存在节点伸缩的情况下，有尽可能多的请求命中原来的机器节点.

### 数据结构
redis 的数据结构包括 list、string、hash、set、zset 五种.它们的底层实现包括 ziplist、hashtable、intset、skiplist 等.
- ziplist: 压缩列表是为了节约内存而开发的，它主要用于 hash 和 list.由一系列特殊编码的内存块构成的列表，一个 ziplist 可以包含多个节点.
- skiplist: 是一种有序的数据结构.一个跳表有多个层 level 组成，通常是 10-20 层，默认是 12 层，每一层都是有序的链表，第 0 层拥有所有数据.它通过在每个节点中维持多个指向其它节点的指针，达到快速访问的目的.它的平均访问效率是 o(log n).

### 数据一致性
但它无法满足数据的强一致性.
**AOF 三个级别**: no 级别 AOF 文件同步由操作系统决定, everysec 每隔一秒执行一次文件同步, always 每次写入文件立即同步. always 十分消耗性能但也并不能满足数据强一致性, 它没有类似 MySQL 的二阶段提交保障写操作中崩溃导致的数据不一致性. 此外文件写入磁盘的过程并非原子操作, 而 MySQl 采用了 double write 解决写入磁盘的一致性.
**主备**: redis 支持主备自动转移故障, 但主从互备也不能保障数据一致性. 当出现同时宕机、网络故障时影响主从同步、过期 key 不主动删除.

## 性能
redis 的单机性能与机器配置有关，一般写性能在 10w qps 左右，读性能为 100w qps.

## 问题
1. redis 的 zset 怎么实现的?(跳表、压缩表、哈希表)
2. 使用 zset 做排行榜时, 如果要实现分数相同时按时间排序怎么实现?
3. binlog 和 redolog 日志?
4. 

## 参考
1. [Redis · 特性分析 · AOF Rewrite 分析](http://mysql.taobao.org/monthly/2016/03/05/)
2. [Redis数据结构底层实现](https://segmentfault.com/a/1190000040206818)
3. [redis中zSet排序原理----skipList跳跃表](https://segmentfault.com/a/1190000022320734)
4. [压缩列表](https://redisbook.readthedocs.io/en/latest/compress-datastruct/ziplist.html)
5. [Redis数据一致性分析](http://baobing.github.io/2017/12/23/Redis/Redis数据一致性分析/)