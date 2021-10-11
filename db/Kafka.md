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
- Consumer Group: 消费者群组，由若干消费者组成的集体。不同的 group 可以对某个 Topic 中的消息进行多次消费，但同一个 Consumer Group 内，对一个消息只能消费一次。

## 可靠性策略


## 参考
1. [kafka 基础知识梳理(一) - 概述](https://www.jianshu.com/p/6b9fa8891026)
2. [Kafka 入门介绍](https://lotabout.me/2018/kafka-introduction/)