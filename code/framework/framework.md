# 简介
golang、c++ 的服务端开发框架，一般需要支持重试、超时处理。

## 基础知识
- 服务雪崩: 在服务链路上，由于服务提供方导致链路上的服务不可用放大，链路上的重试不断放大，导致调用者均不可用。
- 重试机制: 重试 quota，避免雪崩。快速换机重试、服务器过载
- 超时机制: 级联超时、动态超时
- 接入层: 通常是直接面向用户连接或访问的部分
- 逻辑层: 业务逻辑部分
- 微服务: 软件由明确的定义的 API 进行通信的小型独立服务组成

### 微服务开发框架
- RPC 远程调用: 包体序列化、定义传输协议、服务发现
- 同步、异步: 接入层 epoll I/O 多路复用、Work 层同步、异步化网络
- 错误处理: 系统错误、业务错误、网络错误
- 路由: 一致性 hash、屏蔽机器跳过、重试换机
- 过载保护: 根据等待耗时、cpu 等因素拒绝连接访问
- 进程、线程、线程池、协程、协程池
- 协议: rpc/http 协议
- 安全: 参数校验、鉴权
