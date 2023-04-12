# 简介
kafka 是 Apache 基金会开源项目，是一个高吞吐、分布式、分区、多副本、多订阅者的 MQ 系统.

## 基本概念
消息队列一般有两种消费模型，队列模式和发布/订阅模式.
- 队列模式: 多个消费者共同消费一个队列，每条消息只发送给一个消费者.
- 发布/订阅模式: 多个消费者订阅主题，每个消息会发布给所有的消费者.

### 基本组件
- Broker: kafka 集群中一台服务器就是一个 Broker，一个集群由多个 broker 组成，一个 broker 有多个 topic. broker 承担中间缓存和分发的作用, broker 将 producer 发送的数据注册到 consumer 中.
- Topic: 消息主题，一个 topic 即为一个消息队列.消费者可以选择监听指定 topic 的消息队列.
- Producer: 消息生产者，向 topic 发送消息的一方.
- Consumer: 消息消费者，即订阅者，向 topic 中拉取/消费消息的一方
- Consumer Group: 消费者群组，由若干消费者组成的集体.不同的 group 可以对某个 Topic 中的消息进行多次消费，但同一个 Consumer Group 内，对一个消息只能消费一次.消费者组逻辑上是一个订阅者，不同消费者组会收到同一个 topic 的全部消息.而同一个消费者组内的消费者则是以队列模式消费消息的，每条消息只会发送给消费者组队一个消费者.
- Partition: 一个 broker 中可以有多个 topic，一个 topic 可以设置多个 partition(分区). 每个 partition 在物理上对应一个文件夹，存储这个 partition 的所有消息和索引文件.partition 中每个消息分配一个有序的 ID(offset).
- Offset: 偏移 offset 是 partition 分区中消息的有序 ID，它是顺序递增的.通过三元组 <topic,partition,offset> 就可以唯一的定位一条消息.

### Rebalance
Consumer Group 的重平衡 Rebalance 规定额一个 Consumer Group 下的所有 Consumer 如何达成一致，用来分配订阅 Topic 的每个分区.触发 Rebalance 的条件主要有 3 个:
1. Consumer Group 内的成员数量发生变化，比如新的 Consumer 实例加入组或离开组，或者实例崩溃被剔除组.
2. 订阅主题数发生变化.Consumer Group 支持正则表达式的方式订阅主题，若新的主题满足正则表达式导致订阅主题数发生变化.
3. 订阅主题的分区数发生变更.Kafka 当前只能允许增加一个主题的分区数，当分区数增加时，就会触发订阅主题的所有 Group 开启 Rebalance.

在 Rebalance 过程中，所有的 Consumer 实例都会停止消费，等待 Rebalance 完成，这是需要注意的地方，而且 Rebalance 的速度叫慢，曾有上百个 Group 的 Rebalance 成功需几个小时的例子.

### 事务
- At most once(最多一次): 不会重复，但是可能丢失数据.当生产者 ack 超时或者返回错误时，不重试发送消息，会导致消息可能没有写入 kafka topic 中.
- At least once(最少一次): 不会丢失，但是可能导致重复.当生产者的 ack 超时或错误，而此时 broker 已经写入了消息，生产者的重试机制会导致消息被写入两次.
- exactly once(精确一次): 刚好一次，不丢失也不重复，具有幂等性.它需要消息系统本身、生成消息的业务程序及消费消息的业务程序一起完成.

### ack 机制
kafka 的消息确认机制可以选择三种模式，通过设置 acks 参数为 0、-1、1.
- 0: 生产者不会等待 broker 的 ack，这个延迟最低，但是可能在 server 挂掉的时候丢失数据.
- 1: 服务端会等待 ack 值，leader 确认接收到消息后发送 ack，但是 leader 挂掉后不会确保其它副本完成数据复制，可能导致数据丢失.
- -1: 在 1 的基础上，服务端会等待所有 follower 的副本收到数据后才会收到 leader 的 ack ，这样数据不会丢失.

## 特性
kafka 具有一致性和可靠性.
- 一致性: 由于副本并没开放读的能力，所有的操作均有 leader 完成.故 kafka 的一致性保障比较简单，主要指 leader 发生切换前后，通过 HighWatermark 实现.
- 可靠性: 主要通过副本机制实现可靠性，在分区的多个副本中选举一个 leader，读写通过 leader 完成，follower 定期同步 leader 数据，当 leader 挂了之后重新选举新的 leader.通过多副本的冗余数据，确保 kafka 的可靠性，不会轻易丢失数据.在消息确认机制上，可以通过 acks 参数决定不同场景的可靠性保障.

kafka 的主要特性:
- 高吞吐量: 吞吐量达每秒数十万上百万.这么高的吞吐量主要基于顺序读写、零拷贝、分区并发、批量发送、数据压缩等.
- 高并发: 支持上千个客户端(生产者、消费着)读写
- 低延迟: 延迟最低是毫秒级
- 持久性与可靠性: 消息持久化存储与磁盘、支持数据备份
- 集群容错性: 允许 n-1 个节点失败(n 为副本失败)
- 可扩展性: 支持集群动态扩展 

## 监控
监控 kafka，可以从 kafka 主机、 JVM 和 kafka 集群三个维度进行.
- 主机监控: 指的是监控 Broker 所在节点的机器的性能.主要涉及到监控指标包括机器负载、CPU、内存使用率、I/O、TCP 连接数、文件打开数、inode 使用情况等.
- JVM 监控: kafka broker 进程是一个 java 进程，故需要考虑监控 JVM.主要需要关注 FULL GC 的频率和时长、活跃对象大小、应用线程总数等.
- 集群监控: 对 kafka 集群的监控，可以考虑从集群中 broker 进程启动情况、端口建立情况、broker 日志、broker 端线程情况(Log Compaction 线程、副本拉取消息的线程)、关键 JMX 指标、kafka 客户端等.

## 性能指标
**吞吐量**: 单机支持 10w 级别消息传输
**topic 数量**: 几十至几百 topic 时候，吞吐量会大幅下降, 与机器数量有关
**时效性**: 延迟在 ms 级别
**可靠性**: 可以做到 0 丢失

## 使用场景
**异步通信**: 可以让主流程无需等待目标系统响应，将消息异步通知处理目标系统即可. 做到系统快速响应.
**错峰流控与流量削峰**: 在大型活动时，通过队列服务堆积缓存订单等信息，在下游系统有能力处理的时候再处理，避免下游订阅系统因突发流量崩溃.
**日志同步**: kafka 的设计初衷就是为了应对大量日志传输场景，通过异步方式将日志消息同步到消息服务，通过其他组件实时离线分析，对关键日志进行应用监控. 
**延时队列**: 30 分钟未支付自动取消、退款 7 天内未处理自动完成退款、订单 7 天内自动结算等场景.
## 参考
1. [kafka 基础知识梳理(一) - 概述](https://www.jianshu.com/p/6b9fa8891026)
2. [Kafka 入门介绍](https://lotabout.me/2018/kafka-introduction/)
3. [怎么理解 Kafka 消费者与消费组之间的关系?](https://segmentfault.com/a/1190000039125247)
4. [kafka的可靠性与一致性](https://zhuanlan.zhihu.com/p/107705346)
5. [浅谈Kafka特性与架构](https://juejin.cn/post/6844903957664382989)
6. [八年面试生涯，整理了一套Kafka面试题](https://juejin.cn/post/6844903889003610119)
7. [Kafka的Exactly-once语义与事务机制](https://www.cnblogs.com/luxiaoxun/p/13048474.html)
8. [主流MQ的介绍](https://juejin.cn/post/6844904122643120142)
9. [Kafka 为什么会丢消息？](https://www.toutiao.com/article/7149031751204323847/)
10. [基于 Redisson 和 Kafka 的延迟队列设计方案](https://juejin.cn/post/7144969196542099469)