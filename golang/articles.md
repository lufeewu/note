# 简介
一些 golang 相关的文章.

## epoll
常见的 I/O 多路复用有三种 select、poll、epoll. 其中 select 和 poll 每次调用时需要监听所有 fd 集合, 将用户态 fd 拷贝到内核态, 资源耗费大. golang 中的主要使用了 epoll 的三个系统调用结合协程完成异步 I/O. 
- epoll_create: 创建并返回 epfd 句柄.
- epoll_ctl: 向 epfd 中添加、删除、修改监听的 fd.
- epoll_wait: 传入创建返回的 epfd 句柄、超时时间, 返回就绪的 fd 句柄.
<img src="../img/epoll.png">

- eventpoll: 包含 lock、mtx、wq、rdlist 等成员
- fd:
- epfd:
- rdlist:
- et: 边缘触发 edge-triggered
- lt: 水平触发 level-triggered
## 文章
1. <a href="https://medium.com/a-journey-with-go/go-what-does-a-goroutine-switch-actually-involve-394c202dddb7"> Go: What Does a Goroutine Switch Actually Involve? </a>
2. <a href="https://deepu.tech/memory-management-in-golang/">Visualizing memory management in Golang</a>
3. <a href="https://draveness.me/golang/">Go 语言设计与实现</a>
4. <a href="https://jingwei.link/2019/05/26/golang-routine-scheduler.html">Golang 并发问题（五）goroutine 的调度及抢占</a>
5. <a href="https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html">Scheduling In Go : Part II - Go Scheduler</a>
6. <a href="https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-netpoller/">网络轮询器</a>
7. [epoll在Golang的应用](https://zhuanlan.zhihu.com/p/344581947)
8. [如果这篇文章说不清 epoll 的本质，那就过来掐死我吧！](https://www.6aiq.com/article/1564634702930)
9. [linux下epoll如何实现高效处理百万句柄](https://zhuanlan.zhihu.com/p/277664172)

## 书本
1. <a href="https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md">build web application with golang</a>


## 面试
1. <a href="https://github.com/Snailclimb/JavaGuide#高并发">Java学习+面试指南</a>
2. <a href="https://github.com/donnemartin/system-design-primer">The System Design Primer</a>
3. <a href="https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-netpoller/">6.6 网络轮询器</a>
4. <a href="https://github.com/donnemartin/system-design-primer">Go 语言之调度器与 Goroutine</a>
5. <a href="https://mp.weixin.qq.com/s?src=11&timestamp=1587819739&ver=2300&signature=jLRkbkmCkCwT8kGB7Edyv5wGGaRGVYLKGMmmoZ16DOj9iLfro5EW1rHzpiz-nIJa74LpmrUNu6FchioB2ukp8RcXlBsvifXLMezKuX*4dKQ8JqmZBPdIPjkUCBzAOjU2&new=1">最全的常用正则表达式大全</a>
6. <a href="https://www.nowcoder.com/discuss/412272?type=post&order=time&pos=&page=1&channel=">腾讯golang社招面经</a>
7. <a href="https://mp.weixin.qq.com/s?src=11&timestamp=1588230387&ver=2309&signature=TLEMzYVHhbynXH-OvqAso5kwK8c2zGAzk9R0k6qzMVC*ZcuQwh396S32S3cgHw0z2*hlOCc3nxjIOohM3u7TMNFqI1S3wTNEuAGAJFj0WmLCx0UzuAyps2xWoQLd-otV&new=1">Golang面试题</a>

