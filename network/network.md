# 计算机网络
涉及 tcp、http、udp 等.

## OSI 七层网络模型
<img src="../img/osi.png">


## 网络协议

网络协议为计算机网络中进行数据交换而建立的规则、标准或约定的集合.
<img src="../img/protocol.png">
+ 数据传输过程
<img src="../img/network_transfer.png">

+ 异构网络的封包转换问题
+ 网络应用程序视图
    - 数据链路层即 MAC 层，它解决的是局部网络的数据传输能力，如固网、WiFi、3G/4G/5G等
<img src="../img/network_platform.png">


- 1. ARQ 协议是属于哪一层的协议？（链路层？传输层？）
- 2. 数据链路层如何确保帧的顺序？ARQ 协议如何确定重传的帧？

## MQTT 协议
MQTT 全称(Message Queuing Telemetry Transport 消息队列遥测传输)，是一种基于发布/订阅模式的轻量级物联网消息传输协议。

MQTT 传输的消息可以简化为主题(Topic)和载荷(Payload)两部分:
- Topic: 消息主题，订阅者向代理订阅主题后，一旦代理收到相应的主题消息，就会向订阅者转发该消息。
- Paylaod: 消息载荷，订阅者在消息中真正关心的部分，通常是业务相关的。


### MQTT 概念
MQTT 协议包含多个基本概念如下:
- 客户端(Client): 使用 MQTT 的程序或设备，它可以连接到服务端，给其它应用发布消息，订阅接受消息，取消订阅等。
- 服务器(Server): 在客户端之间充当中介角色的程序或设备，他可以接受客户端连接，接受客户端发布的消息，处理客户端的订阅和取消订阅，转发消息等。
- 会话(Session): 每个客户端与服务器建立连接后就是一个会话，客户端与服务器之间有状态的交互。会话可以存在于一个网络连接之间，也可以跨越多个网络连接。
- 订阅(Subscription): 包含主题过滤器(Topic Filter)和最大服务质量(QoS)等级。订阅与单个会话关联，会话可以包含多余一个的订阅。
- 主题名(Topic Name): 附加在应用消息上的一个标签，被用于匹配服务端已存在的订阅。服务端会向所有匹配订阅的客户端发送消息。
- 主题过滤器(Topic Filter): 仅在订阅时使用的主题表达式，可以包含通配符，以匹配多个主题名。
- 载荷(Payload): 对于 PUBLISH 报文来说载荷就是业务消息，它可以是任意格式(二进制、十六进制、普通字符串、JSON 字符串、Base64等)的数据。

此外，MQTT 的概念还有服务质量(QoS)、清除会话(Clean Session)、保活心跳(Keep Alive)、保留消息(Retained Message)、遗嘱消息(Will Message)等.
- 服务质量(QoS): MQTT 消息提供了三种消息服务质量等级，保证在不同网络环境下的消息传递可靠性。QoS 0 消息最多传递一次，客户端不可用则会丢失消息。QoS 1 消息至少传递一次，QoS 2 消息仅传递一次。
- 清除会话(Clean Session): 客户端发起 CONNECT 请求时，可以通过 Clean Session 标志设置是否创建全新会话。
- 保活心跳(Keep Alive): 发起 CONNECT 请求时，可以通过 Keep Alive 参数设置保活周期。在客户端没有报文发送时，会定时发送 2 字节的 PINGREQ 心跳报文。服务端在 1.5 个 Keep Alive 周期内，既没有收到客户端的发布订阅报文，也没有收到 PINGREQ 心跳报文，则断开客户端连接。
- 保留消息(Retained Message): MQTT 客户端向服务器发布消息时，可以设置保留消息标志，保留消息会驻留在消息服务器，后来的订阅者订阅主题时可以接受到最新的一条保留消息。
- 遗嘱消息(Will Message): MQTT 客户端向服务端发送 CONNECT 请求时，可以携带遗嘱消息。MQTT 客户端异常下线时（客户端断开前未向服务器发送 DISCONNECT 消息)，MQTT 消息服务器会发布遗嘱消息。


## 问题
1. 10m buffer 里面存满数据, 将数据尽量发出去, 允许部分丢包, 使用 tcp 好还是 udp 好?
    - udp 可以保证速度、没有重传机制、没有阻塞机制、速度最快
    - tcp 可以保证尽量不丢包、丢包有重传, 网络环境差的时候推荐 tcp

## 参考
1. [MQTT 协议 10 分钟快速入门](https://www.emqx.com/zh/blog/get-started-with-mqtt-in-ten-mins)
2. [](https://note.grianchan.com/网络/网络.html)