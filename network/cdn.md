# 简介
CDN 全称 Content Delivery Network, 即内容分发网络. 基本思路是尽可能避开互联网上影响数据传输速度和稳定性的平

## 基础概念
- 回源: 指通过客户端请求访问资源时, 如果 CDN 节点上未缓存该资源, 或者您部署预热任务给 CDN 节点时, CDN 节点会回源站获取资源.
- 回源率: 包括回源请求数、回源流量比两种.
- CDN 多级缓存: 浏览器本地缓存、CDN 边缘缓存.
- CDN 缓存策略: 边缘节点缓存策略因服务商不同而不同, 一般通过 http 中的 Cache-control 控制: max-age 设置缓存时间. 客户端请求 CDN 节点时, 首先判断缓存数据是否过期, 若没有过期, 则直接将缓存数据返回给客户端. 否则, CDN 节点向源站发出请求, 从源站拉取最新数据, 更新本地缓存, 并将最新数据返回给客户端. CDN 服务商一般提供基于文件后缀、目录等多个维度指定 CDN 缓存时间，提供更精细化的缓存管理. CDN 缓存时间将会影响到回源率. 缓存时间短, 将会频繁回源. 缓存时间长会导致更新时间慢.
- CDN 层级划分:
    + 边缘层: CDN 系统中直接面向用户负责给用户提供内容服务的 Cache 设备部署在整个 CDN 网络的边缘位置, 将这一层成为边缘层.
    + 中心层: 负责全局的管理和控制, 也保存了最多的内容 cache. 在边缘设备未能命中 Cache 时, 需要向中心层设备请求. 中心层未能命中时, 需要向源站请求. 有的中心层设计具备向用户提供服务能力, 有的则只向下一层提供服务.
    + 区域层: 若 CDN 系统比较庞大, 边缘层向中心层请求内容太多, 会造成中心层负载压力太大. 此时, 需要在中心层和边缘层之间部署一个区域层, 负责区域的管理和控制, 也能提供一些内容 Cache 供边缘层访问.

## 参考文献
1. [内容分发网络](https://baike.baidu.com/item/内容分发网络/4034265)
2. [这就是CDN回源原理和CDN多级缓存啊!](https://cloud.tencent.com/developer/article/1439913)