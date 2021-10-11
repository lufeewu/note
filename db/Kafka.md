# 简介
kafka 是 Apache 基金会开源项目，是一个高吞吐、分布式、分区、多副本、多订阅者的 MQ 系统。

## 基本概念
消息队列一般有两种消费模型，队列模式和发布/订阅模式。
- 队列模式: 多个消费者共同消费一个队列，每条消息只发送给一个消费者。
- 发布/订阅模式: 多个消费者订阅主题，每个消息会发布给所有的消费者。

### 基本组件
- Broker: kafka 集群中一台服务器就是一个 Broker，一个集群由多个 broker 组成，一个 broker 有多个 topic. broker 承担中间缓存和分发的作用, broker 将 producer 发送的数据注册到 consumer 中.
- Topic: 消息主题，一个 topic 即为一个消息队列。消费者可以选择监听指定 topic 的消息队列。
- Producer: 消息生产者，向 topic 发送消息的一方。
- Consumer: 消息消费者，即订阅者，向 topic 中拉取/消费消息的一方
- Consumer Group: 消费者群组，由若干消费者组成的集体。不同的 group 可以对某个 Topic 中的消息进行多次消费，但同一个 Consumer Group 内，对一个消息只能消费一次。消费者组逻辑上是一个订阅者，不同消费者组会收到同一个 topic 的全部消息。而同一个消费者组内的消费者则是以队列模式消费消息的，每条消息只会发送给消费者组队一个消费者。
- Partition: 一个 broker 中可以有多个 topic，一个 topic 可以设置多个 partition(分区). 每个 partition 在物理上对应一个文件夹，存储这个 partition 的所有消息和索引文件。partition 中每个消息分配一个有序的 ID(offset).
- Offset: 偏移 offset 是 partition 分区中消息的有序 ID，它是顺序递增的。通过三元组 <topic,partition,offset> 就可以唯一的定位一条消息。

## 特性
kafka 具有一致性和可靠性。
- 一致性: 由于副本并没开放读的能力，所有的操作均有 leader 完成。故 kafka 的一致性保障比较简单，主要指 leader 发生切换前后，通过 HighWatermark 实现。
- 可靠性: 主要通过副本机制实现可靠性，在分区的多个副本中选举一个 leader，读写通过 leader 完成，follower 定期同步 leader 数据，当 leader 挂了之后重新选举新的 leader。通过多副本的冗余数据，确保 kafka 的可靠性，不会轻易丢失数据。在消息确认机制上，可以通过 acks 参数决定不同场景的可靠性保障。

kafka 的主要特性:
- 高吞吐量: 吞吐量达每秒数十万上百万。这么高的吞吐量主要基于顺序读写、零拷贝、分区并发、批量发送、数据压缩等。
- 高并发: 支持上千个客户端(生产者、消费着)读写
- 低延迟: 延迟最低是毫秒级
- 持久性与可靠性: 消息持久化存储与磁盘、支持数据备份
- 集群容错性: 允许 n-1 个节点失败(n 为副本失败)
- 可扩展性: 支持集群动态扩展 


## 性能指标


## 参考
1. [kafka 基础知识梳理(一) - 概述](https://www.jianshu.com/p/6b9fa8891026)
2. [Kafka 入门介绍](https://lotabout.me/2018/kafka-introduction/)
3. [怎么理解 Kafka 消费者与消费组之间的关系?](https://segmentfault.com/a/1190000039125247)
4. [kafka的可靠性与一致性](https://zhuanlan.zhihu.com/p/107705346)
5. [浅谈Kafka特性与架构](https://juejin.cn/post/6844903957664382989)