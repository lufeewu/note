# 简介
sync 是 golang 的基础库之一，它提供基本的同步原语，如互斥锁. 除了 Once 和
WaitGroup 类型之外，其它大多数都是提供低级别例程使用. 更高级别的同步最好通过通道及通信完成.

## sync

## chan

## sync + chan 进行并发及数量控制
通过 sync 可以并发开启协程，保证主程序在协程都处理完成后推出. 通过 chanel 可以控制并发的数量，限制 goroutine 数量过多. 下面的程序是一个简单的用 sync+chan 进行并发及 goroutine 数量限制的方案.

    func syncChan() {
        wg := sync.WaitGroup{}
        ch := make(chan int, 100)
        for i := 0; i < 1203; i++ {
            ch <- 1
            wg.Add(1)
            go func(i int) {
                defer func() {
                    <-ch
                    logrus.Infoln("over", i)
                    wg.Done()
                }()
                time.Sleep(1 * time.Second)
            }(i)
        }

        wg.Wait()
    }
