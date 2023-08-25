# 简介
Memcached 是一个基于 key-value 存储, 用于存储小块数据. 它是一套分布式的告诉缓存系统.

## slab allocation 
slab allocation 是 Memcached 的内存分配机制, 它以 slabs 为单位, 根据初始 chunk 大小、增长因子、存储数据的大小实际划分出多个不同的 slabs class, slab class 中包含若干个等大小的 trunk 和一个固定 48 byte 的 item 信息. trunk 则是按页存储, 每一页为一个 Page.
- Page: 分配给 Slab 的内存空间, 默认为 1MB, 分配后就得到一个 Slab, Slab 分配之后内存按照固定字节大小等分为 chunk.
- Chunk: 用于缓存记录 k/v 值的内存空间. Memcached 会根据数据大小选择存到哪一个 chunk 中, 比如有 128 bytes、64 bytes 等多种.
- Slab Class: Slab 按照 Chunk 的大小分组, 组成不同的 Slab Class, 第一个 Chunk 大小为 96B 的 Slab 为 Class 1, Chunk 120B 为 Class 2, 

## 集群
Memcached 是一个单进程多线程模型的缓存数据库, 它的集群主要体现在数据的分片中, 是靠客户端实现的, 它的服务端之间没有通信, 它通过客户端实现的分布式算法把数据保存到不同的 Memcahced 服务端中. 常见的分布式算法有:
- 取模: 余数计算的方法比较简答, 数据的分散性也好. 但是添加或移除服务器时, 缓存重组的代价较大.
- 一致性 hash: 通过将 Memcached 服务器节点的 hash 值映射到环上, 由 key 的映射匹配到指定的服务器.

## 参考
1. [memcached介绍与它的工作原理](https://blog.51cto.com/u_15105742/5282807)
2. [内存分配机制Slab Allocation](https://www.cnblogs.com/douJiangYouTiao888/p/6267569.html)
3. [memcached分布式原理与实现](https://juejin.cn/post/6844903665875025933)