# 简介
http 协议是互联网最广泛的应用层协议，它也是经过了一段时间的发展，它有三个重要阶段 1.0、1.x、2.0.


## http 1.0 
http 1.0 是比较久远的版本，每次请求都需要建立三次握手连接，性能较差。

## http 1.x
http 1.x 在 http 1.0 的基础上增加了 pipline，多个连续的请求可以不用每次都建立三次握手，大大减少了网络延迟。

## http 2.0
http 2.0 在速度上有了质的飞升，它被用于 gRPC 中。它主要进行了二进制分帧、多路复用、头部压缩、服务器推送等方式进一步提升了 http 1.x 的性能。

## https
https(HyperText Transfer Protocol Secure) 是安全通信的 http 协议，基于 SSL/TLS 进行安全加密。

## 参考
1. [http2.0为什么那么快](https://zhuanlan.zhihu.com/p/380933480)