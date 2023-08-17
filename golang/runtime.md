# 简介
golang 实现了自己的 runtime，功能包括 GC、goroutine、内存分配等. 极大的提高了 go 语言的效率.

## tight loop
golang 1.14 版本之前，对于 tight loop 不能被 gc 中断. 对于小的 tight loop 的也并没有进行并发处理. 下述的代码中 A 的输出和 B 的输出并非交叉而是连续的.

        var wg1 sync.WaitGroup
        func a() {
            defer wg1.Done()
            for i := 0; i < 10; i++ {
                fmt.Println("A:", i)
            }
        }
        func b() {
            defer wg1.Done()
            for i := 0; i < 10; i++ {
                fmt.Println("B:", i)
            }
        }
        func testB() {
            // fmt.Println(runtime.NumCPU())
            runtime.GOMAXPROCS(12)
            wg1.Add(2)
            go a()
            go b()
            wg1.Wait()
        }

## runtime 库
- NumCPU() 
- GOMAXPROCS()
- Gosched()
- Goexit()
- NumGoroutine()
- GOOS

## 并发
n++ 并发安全问题

    func race() {
        num := 1000
        wait := make(chan struct{}, num)
        n := 0
        for i := 0; i < num; i++ {
            go func() {
                // 译注：注意下面这一行
                n++ // 一次访问: 读, 递增, 写
                wait <- struct{}{}
            }()
        }

        // 译注：注意下面这一行
        // n++ // 另一次冲突的访问
        for i := 0; i < num; i++ {
            <-wait
        }
        fmt.Println(n) // 输出：未指定
    }
### 协程
Coroutine 轻量级线程, 它的切换完全在用户态进行, 相比线程、进程效率更高. 

- 进程控制原语: 建立、撤销、等待、唤醒
- 进程状态: D(TASK_UNINTERRUPTIBLE) 不可中断睡眠状态、R(TASK_RUNNING) 可执行状态、S(TASK_INTERRUPTIBLE) 可中断的睡眠状态、T/t(TASK_STOPPED or TASK_TRACED) 暂停状态或跟踪状态、X(TASK_DEAD - EXIT_DEAD) 退出状态，进程即将被销毁、Z(TASK_DEAD - EXIT_ZOMBIE) 退出状态，进程成为僵尸进程

## pprof
gcBgMarkWorker
mallocgc

## 数据结构
golang 有丰富的数据结构，除了基础的 int、string、char、byte 等，还提供了 map、slice、channel 等类型。

### map 底层
golang 的 map 主要基于哈希表原理，能够实现 o(1) 时间复杂度的操作。
- 内存分配: 当通过 make 申请 map 时，当指定的 hint 小于等于 8 时，直接在栈上分配一个 bucket(每个 bucket 可以存储 8 对 kv). 而当 hint 大于 8 小于等于 52 时，会在堆上分配 bucket，但不会分配 overflow bucket。当 make 的 hint 大于 52 时，会在堆上分配 bucket 和 extra 的 overflow bucket。
- buckets: 一段连续空间 2^B 大小的 bucket 数组。当触发 buckets 扩容，则会增长 2 倍大小。
- oldbuckets: 用于实现增量扩容，若扩容正在进行中，则 oldbuckets 是有值的。
- mapextra: mapextra 中包含 overflow 的 bucket(也氛围 overflow 和 oldoverflow) 。 overflow bucket 用于在出现哈希冲突时进行存储冲突 kv，即拉链法解决冲突。
- 寻址: 通过将 key 进行 hash，并将 hash 所得的低位值作为 buckets 数组的索引，然后高位 hash 比较 bucket 中的 hash 是否一致，一致则在 bucket 中寻找指定 key。不一致则继续在 overflow buckets 中寻找。

## 内存分配
golang runtime 的内存分配算法主要源于 c 语言开发的 TCMalloc 算法, 全称 Thread-Caching Malloc, 核心思想是把内存分为多级管理, 从而降低锁的粒度.
- 内存管理: 在启动时, 向操作系统申请一块虚拟地址空间的内存, 切成小块进行管理. 内存快被分配成如下三个区域.
    + arena 区域: 堆区, 大小为 512GB. go 动态分配的内存都在这个区域, 它把内存分割成 8KB 大小的页, 一些页组合起来称为 mspan.
    + bitmap 区域: 区域大小为 16GB, 标识 arena 区域哪些地址保存了对象, 并且用 4bit 标志位表示对象是否包含指针、GC 标记信息.
    + spans 区域: 区域大小为 512 MB, 存放 span 的指针, 每个指针对应一页.
- 内存分配: 由内存分配器完成, 由 3 种组件构成.
    + mcache: 每个工作线程都会绑定一个 mcache, 本地缓存可用 mspan 资源.
    + mcentral: 为所有 mcache 提供切分好的 mspan 资源. 每个 central 保存特定大小的全局 mspan 列表, 包括分配和未分配的.
    + mheap: 代表 go 程序特有的所有堆空间, go 程序使用一个 mheap 的全局对象 _mheap 来管理内存.
- 分配流程: 内存分配对象时, 根据对象大小分成三类, 小于等于 16B、一般对象(大于 16B, 小于等于 32KB)、大对象(32KB). 大体流程如下:
    1. 32Kb 对象直接在 mheap 上分配
    2. <= 16B 的对象使用 mcache 的 tiny 分配器分配.
    3. 大于 16B 小于等于 32KB 的对象, 先计算对象的规格大小, 然后使用 mcache 中相应规格大小的 mspan 分配.
        + 如果 mcache 没有相应规格大小的 mspan, 则向 mcentral 申请
        + 如果 mcentral 没有相应规格大小的 mspan, 则向 mheap 申请
        + 如果 mheap 中没有合适大小的 mspan, 则向操作系统申请


## 参考
1. [golang map底层实现](http://yangxikun.github.io/golang/2019/10/07/golang-map.html)
2. [解剖Go语言map底层实现](https://studygolang.com/articles/14583)
3. [深入解析 go - 2.3 map的实现](https://tiancaiamao.gitbooks.io/go-internals/content/zh/02.3.html)
4. [Linux系统之进程状态](https://cloud.tencent.com/developer/article/1568077)
5. [深入golang runtime的调度](https://zboya.github.io/post/go_scheduler/#go进程的启动)
6. [图解Go语言内存分配](https://juejin.cn/post/6844903795739082760)
7. [Golang 内存组件之mspan、mcache、mcentral 和 mheap 数据结构](https://segmentfault.com/a/1190000039815122)