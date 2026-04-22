## 简介
机器人相关技术.

## DDS
数据分发服务(Data Distribution Service, DDS) 是一种用于分布式实时通信的"以数据为中心"的中间件协议. 它采用发布/订阅模型, 广泛应用于汽车智驾、机器人 (ROS 2)、工业自动化等领域, 以保证数据传输的实时性、高可靠性和灵活性.
- Pub/Sub 架构: 发布者 Pub 向网络中发布特定主题(topic)的数据、订阅者 Sub 声明感兴趣的主题后自动接受相关数据.
- RTPS 协议: Real-Time Publish-Subscribe 协议, 定义了数据序列化、发现机制、传输格式等, 支持 UDP/IP、共享内存、自定义传输层.

### DDS 与 Kafka 发布-订阅（Pub/Sub）模型对比

| **对比维度**         | **DDS (Data Distribution Service)**                                                                 | **Kafka (Apache Kafka)**                                                                                 |
|----------------------|-----------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| **核心定位**         | 实时数据总线（Real-time Data Bus）<br>以数据为中心，分发“当前系统状态”                                | 分布式日志流平台（Distributed Log Streaming Platform）<br>以消息为中心，持久化“完整事件序列”             |
| **通信架构**         | **去中心化**（Peer-to-Peer）<br>无 Broker，参与者通过 RTPS 协议直接通信                               | **中心化**（Broker-based）<br>Producer → Kafka Broker → Consumer                                         |
| **发现机制**         | **自动动态发现**<br>基于 Topic + QoS 自动匹配 Publisher/Subscriber                                   | **静态配置**<br>Consumer 必须显式订阅已知 Topic 名称                                                     |
| **Topic 语义**       | **强类型**<br>由 IDL 定义结构，支持内容过滤（Content Filtered Topic）                                | **弱类型 / 字节流**<br>Topic 为字符串名，消息为字节数组（需应用解析）                                    |
| **数据模型**         | **状态导向（State-Oriented）**<br>关注最新值或当前有效状态                                           | **事件导向（Event-Oriented）**<br>关注不可变的事件日志                                                   |
| **历史回溯**         | ❌ 默认不支持<br>✅ 可通过 `Durability=TRANSIENT_LOCAL` 获取有限历史（由 Writer 缓存）                | ✅ 原生支持<br>消息持久化到磁盘，Consumer 可从任意 offset 回放                                           |
| **Offset 机制**      | ❌ **无全局 offset**<br>内部有序列号但不暴露给应用层                                                 | ✅ **核心概念：offset**<br>每个 Partition 有单调递增 offset，可提交消费位置                              |
| **读取语义**         | `read()`（保留缓存） / `take()`（移除缓存）<br>基于 `History QoS` 窗口                              | `poll()` 拉取<br>基于 offset 范围                                                                       |
| **QoS 支持**         | ✅ **高度灵活**<br>20+ 种策略：<br>- Reliability（可靠/尽力）<br>- Durability（持久性）<br>- Deadline（截止时间）<br>- Lifespan（生命周期）<br>- Ownership（所有权）等 | ⚠️ **有限配置**<br>主要通过：<br>- `acks`<br>- `replication.factor`<br>- `min.insync.replicas` 控制可靠性 |
| **端到端延迟**       | **极低**（微秒 ~ 毫秒级）<br>适合硬实时系统（如控制指令）                                            | **较低**（毫秒 ~ 秒级）<br>不适合严格实时场景                                                            |
| **吞吐量**           | 中~高（万级 msg/s）                                                                                | **极高**（百万级 msg/s）                                                                                |
| **持久化**           | ❌ 默认内存缓存<br>✅ 可选 `PERSISTENT`（需插件）                                                    | ✅ **原生存储**<br>消息写入磁盘，支持多副本                                                              |
| **部署复杂度**       | 无中心节点，嵌入式友好<br>但 QoS 配置复杂                                                           | 需维护 Broker + ZooKeeper/KRaft 集群<br>运维复杂                                                         |
| **典型应用场景**     | - 自动驾驶<br>- 工业控制（PLC/机器人）<br>- 航空航天<br>- 医疗设备<br>- 实时仿真                    | - 日志收集与分析<br>- 事件溯源<br>- 微服务异步通信<br>- 大数据管道（ETL）<br>- 用户行为追踪            |

**对比总结**
- **选 DDS 当**：你需要 **超低延迟、确定性通信、复杂 QoS 控制**，且系统运行在 **资源受限的边缘/嵌入式环境**。
- **选 Kafka 当**：你需要 **高吞吐、消息持久化、历史回溯、大规模数据集成**，且能接受 **毫秒级延迟**。
> 💡 **混合架构趋势**：  
> 在智能汽车、工业物联网中，常见 **DDS + Kafka 混合架构**：  
> - **车内/设备内**：用 **DDS** 实现实时控制（传感器 → 控制器）  
> - **车云/边缘云**：用 **Kafka** 上传遥测数据供大数据分析


##