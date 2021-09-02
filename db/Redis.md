# 简介
Redis 的知识

## 知识点
关于 Redis，需要掌握基本数据结构、跳表、Hash 冲突解决、持久化、主从、分布式、性能等知识。
### 持久化
redis 的持久化方式主要是 AOF、RDB 。
AOF: 保存 binlog，将所有写入命令及参数写入文件保存在磁盘。
- AOF Rewrite 机制: 随着时间的推移，AOF 文件会膨胀的比较厉害，频繁的写入 AOF 文件将影响性能。系统通过 AOF Rewrite 机制定期重写 AOF 文件，将数据库的数据以协议方式保存到新的 AOF 文件中，重写的 AOF 小于原 AOF 文件，从而达到减小 AOF 文件目的。但是 AOF 也会占用较多的内存，写入 AOF buffer 过程中会阻塞工作线程。
- REWRITE: 主线程中重写 AOF，阻塞工作线程，生产环境中几乎不使用。
- BGREWRITE: 在后台子进程中重写 AOF，不阻塞工作线程。
RDB: 将 Redis 的当前数据以快照二进制的方式保存在磁盘，通过 RDB 的快照恢复数据更快。

### 主从
Redis 的高可用灾备方案主要是主从方式。
- 一主一从:
- 多主多从:
- 分片:


## 参考
1. [Redis · 特性分析 · AOF Rewrite 分析](http://mysql.taobao.org/monthly/2016/03/05/)