# vendor
用于支持本地包管理依赖，通过 vendor.json 文件记录依赖包版本，可以将项目依赖的外部包拷贝到项目下的 vendor 目录下。
### 命令
govendor init

govendor add +external

govendor list

govendor list -v fmt

govendor fetch golang.org/x/net/context@{version-id}

govendor fetch golang.org/x/net/context@v1

govendor fetch golang.org/x/net/context@=v1

govendor fetch golang.org/x/net/context

govendor fmt +local

govendor install +local

govendor test +local
  
# dep

### 命令
dep init
![lock toml vendor 关系](./img/28968009-f49a4a6a-78eb-11e7-93cf-e695d45488da.14d8c0f3.png)
dep ensure

dep ensure -add 

dep check

dep status 

# 语言特性
1. 基础
+ 工作区和GOPATH
+ 命令源码文件、库源码文件
+ 25 个关键字不允许定义
   + var、const、**package**、import
   + func、return、**interface**、struct、type、**map**、range
   + go、select、**chan**
   + if、else、switch、case、default、fallthrough、for、break、continue
   + goto、**defer**
+ 预定义
    + 内建常量 true、false、iota、nil
    + 内建类型 
        - int、int8、int16、int32、int64
        - uint、uint8、uint16、uint32、uint64、uintprt
        - float32、float64、complex64、complex128
        - bool、byte、rune、string、error
    + 内建函数
        - make、len、cap、new、append、copy、close、delete
        - complex、real、imag
        - **panic**、**recover**
+ 变量、常量、函数、结构体、接口
    + **浅拷贝，值类型和引用类型**
    + 指针、uintptr、unsafe.Pointer、不可寻址值
    + 嵌入字段
    + 接口类型、动态类型、静态类型、**接口赋值**、**零值**、iface、无类型 nil
    + 常量默认类型
    + byte、uint8
    + value.(type)
    + **副本**
+ 类型推断、代码块、类型断言
+ 数组、切片
    + 切片的容量增长
+ container/list 包,指针
+ map(字典、键值对、键元素对、映射)
+ 通道 channel
    + make、<-、select case
    + FIFO 队列
    + 副本、复制
    + 缓冲通道、非缓冲通道
    + 单向通道、双向通道
    + 同步、异步
+ goroutine 协程
    + go 语句 
    + sync.Pool、sync.Map、sync.WaitGroup
    + CSP 并发模型
    + GPM (goroutine、processor、machine) 
    + runtime.GOMAXPROCS(maxProcs)
    + goroutine pool
    + struct{}{}
    + sync/atomic
    + > 怎样让我们启用的多个 goroutine 按照既定的顺序运行?
+ if、for、switch 语句
    + range、
![golang GMP 模型](./img/golang_GMP.png)

+ **panic、recover、defer**
+ package
    + **error**
        + 卫述语句
        + error type
        + 错误处理
        + 立体的错误类型体系
        + 扁平的错误值列表
    + fmt
    + io.Writer
    + bufio
    + image

+ panic、defer、recover
    + panic
        + 场景系统 panic 原因: 数组越界、空指针引用、断言失败、map 操作错误、除数为 0、调用未实现的方法、通道 chan 操作错误、goroutine 竞争资源导致死锁、非线程安全操作如 map、内存不足
        + 数组越界 index out of rang
        + runtime error: index out of ra...
        + exit status 2
        + 从 panic 被引发到程序终止运行的大致过程是什么？
            1. 建立 panic 详情
            2. 沿着调用栈的反方向传播至顶端, main 函数
            3. go runtime 回收，程序崩溃
            4. 打印 panic 详情
        + 意外 panic, 主动 panic()函数
    + recover
    + deferc
        + 不支持 go 语言内建函数调用
        + 不支持 unsafe 包中的函数的调用表达式
        + 倒序执行、FILO 队列（栈）
        + 可以在 defer 中引发 panic 么？

+ 测试
    + go 程序测试
        + 功能测试（test）
        + 基准测试（benchmark, 性能测试）
        <img src="./img/test_benchmark.png">
        + 示例测试（example）
        + go clean -cache
        + -cpu、-count、-parallel、-benchmark、-benchtime
        + GPM(Goroutine、Processor、Machine)
    + 程序监测
    + go 语言标准库代码用法
    + 单元测试、API 测试、集成测试、灰度测试

+ 互斥锁 sync.Mutex、sync.
    + 互斥锁保证任何时刻只有一个 goroutine 可以访问共享资源。读写锁则允许多个 goroutine 同时读取共享资源，写操作是互斥的。
    + 竞态条件 (race condition)
    + 同步、临界区（critical section）、互斥量（mutual exclusion）
    + lock、unlock、deadlock
    + time.Ticker
    + 共享资源（存储、计算、I/O、网络等）
    + 互斥锁和读写锁的指针类型都实现了哪一个接口？（Lock接口）
    + 怎样获取读写锁中的读锁？（Rlock）

+ 条件变量 sync.Cond
    + 等待通知(wait)，wait 方法做了什么？
    + 单发通知(signal)
    + 广播通知(broadcast)
    + sendCond、recvCond
    + *sync.Cond类型的值可以被传递吗？那sync.Cond...

+ 原子操作 (sync/atomic)
    + sync/atomic 包提供了几种原子操作？可操作的数据类型又有哪些？
    + atomic.Value
    + 加法（Add）、比较并交换（compare and swap，简称 CAS）、加载（load）、存储（load）和交换（swap）
    + 原子操作减法（无符号数的减法）
    + CAS 比较交换操作相比交换操作有什么优势？
    + 自旋锁（spinlock）

+ sync.WaitGroup 和 sync.Once
    + sync.WaitGroup、Add、Done、Wait
    + GoF 单例模式
    + 在使用 waitgroup 值实现一对多的 goroutine 协作流程时，怎样才能让分发子任务的 goroutine 获得各个子任务的具体执行结果呢？

+ context.Context
    + context.Background、context.WithCancel、context.CancelFunc、context.WithDeadline、context.WithTimeout
    + <img src='img/golang_context.png'>
    + context 值在传达撤销信号的时候是广度优先的，还是深度优先的？其优势和劣势都是什么？

+ 临时对象池 sync.Pool
    + Put、Get
    + pp 类型
    + 临时池存储值所用的数据结构是怎样的？
    + 临时对象池是怎样利用内部数据结构来取值的?
    + runtime.GC
+ sync.Map 并发安全字典
    + map 并发安全
    + 并发安全字典对健的类型有要求么？（不支持函数类型、字典类型、切片类型）
    + 怎样保证并发安全字典中的键值类型正确性？
        + 让并发安全字典只能存储某个特定类型的键
        + 接受动态的类型设置，并在程序运行的时候通过反射操作进行检查（reflect.Type)
    + sync.Map dirty 字段

+ Unicode 
    + Ascii、Unicode、UTF-8、byte（http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html）
    + rune
    + 使用带有 range 子句的 for 语句遍历字符串的时候应该注意什么？
    + 判断一个 Unicode 字符是否为单字节有几种方式？

+ strings 包与字符串操作
    + strings.Builder、strings.Reader
    + 裁剪、拼接（切片、+）
    + strings.Builder 类型在使用上有约束么？
    + 为什么说 strings.Reader 类型的值可以高效地读取字符串？(已读计数)
    + *strings.Builder 和 *strings.Reader 分别实现了哪些接口？这样做有什么好处？

+ bytes 包与字符串操作
    + Buffer 长度、总长度, Len、Cap
    + 已读计数
    + Truncate 
    + 扩容策略是怎样的？
    + 哪些方法可能造成内容泄露？
    + 深度拷贝、副本传递避免泄露
    + strings.Builder 和 bytes.Buffer 的 string 方法，哪一个高效？

+ **I/O**
    > io 包中共有简单接口 11 个. 其中读取操作相关的接口 5 个，写入相关的接口 4 个，关闭操作相关的 1 个以及 1 个读写位置设定相关的接口
    + strings.Builder 指针类型实现接口包括 io.Writer、io.ByteWriter、fmt.Stringer、io.stringWriter
    + strings.Reader 指针类型实现接口
        - io.Reader
        - io.ReaderAt
        - io.ByteReader
        - io.RuneReader
        - io.Seeker
        - io.ByteScanner
        - io.RuneScanner
        - io.WriterTo
    + bytes.Buffer 指针类型实现接口
        - io.Reader
        - io.ByteReader
        - io.RuneReader
        - io.ByteScanner
        - io.RuneScanner
        - io.WriterTo
        - io.Writer
        - io.ByteWriter
        - io.ReaderFrom
    + io.Reader 扩展接口
        - io.ReadWriter
        - io.ReadCloser
        - io.ReadWriteCloser
        - io.ReadSeeker
        - io.ReadWriteSeeker
    + io.Reader 实现接口
        - io.LimitedReader
        - io.SectionReader
        - io.teeReader
        - io.multiReader
        - io.pipe
        - io.PipeReader
    + io 包中的接口都有哪些？它们之间都有着怎样的关系？
        - 核心接口 io.Reader、io.Writer、io.Closer
        - 四大操作类,读取、写入、关闭、读写位置设定
    + io 包中的同步内存管道的运作机制是是么？
    <img src='img/io_interface.png'>

+ bufio ( buffered I/O , 内置缓冲区)
    + Reader
        - 字段 buf、rd、r、w、err、lastByte、lastRuneSize
        - bufio.Reader 类型的读取方法有哪些不同？(4个读取流程代表)
    + Scanner
    + Writer 和 ReadWriter
    + Flush 方法
    + 内容泄露 Reader、Peek、ReadSlice、ReadLine 

+ os 包中的 API
    > APi 基于操作系统，为使用操作系统功能提供高层次支持，同时不依赖具体的操作系统
    + 帮助使用操作系统中的文件系统、权限系统、环境变量、系统进程及系统信号
    + dir、env、error、exec、export、executable、file、getwd、path、pipe、removeall、stat、sys、types、wait、signal、user
    + os.File
        - 类 Unix 的一切都可以看做是文件
        - 文本文件、二进制文件、压缩文件、目录
        - 符号链接、物理设备、命名管道、套接字（socket）
        - os.File 类型实现了哪些 io 包中的接口？
        - 怎样获得一个 os.File 类型的指针值？（Create、NewFile、Open、OpenFile）
        - os 错误值 os.PathError、os.ErrInvalid、os.ErrPermission、os.ErrExist
    + 可以用于 File 值的操作模式都有哪些？（只读、读写、只写）
    + 怎样设定常规文件的访问权限？
    + 怎样通过 os 包中的 API 创建和操作一个系统进程？

+ 网络
    + socket、IPC（Inter-Process Communication）
        - 系统信号（signal）
        - 管道（pipe）
        - 套接字（socket）
        - 文件锁（file lock）
        - 消息队列（message queue）
        - 信号灯（semaphore，信号量）
    + syscall
        - syscall.Socket, 3 个参数表示通信域、类型及使用协议
        <img src="img/socket.png">
    + net
        - net.Dial(network, address string)
        - network 可选 9 个值: tcp、tcp4、tcp6、udp、udp4、udp6、unix、unixgram、unixpacket
        - 调用 net.DialTimeout 函数时超过给定的时间意味着什么？
        - net.Dialer 类型 
        - DialContext 方法
        - 怎样在 net.Conn 类型的值上正确地设定针对读操作和写操作的超时时间？
    + net/http
        - http.Get()、*http.Client
        - http.Client 类型中的 Transport 字段代表着什么？ 
        - *http.Transport、http.RoundTripper
            - IdleConnTimeout
            - DefaultTransport
            - ResponseHeaderTimeout
            - ExpectContinueTimeout
            - TLSHandshakeTimeout
        - http.Server 类型的 ListenAndServe 方法都做了哪些事情？(服务端)
        - 怎样优雅的停止 HTTP 协议的网络服务程序？

+ 性能分析
    + runtime/pprof
    + net/http/pprof
    + runtime/trace
    + 概要文件解析工具 go tool pprof 、go tool trace
    + 概要文件: CPU 概要文件、内存概要文件、阻塞概要文件
    + go test
    + 怎样让程序对 CPU 概要信息进行采样？（runtime/pprof）
        - runtime/pprof.StartCPUProfile
    + 怎样设定内存概要信息的采样频率？（runtime.MemProfileRate)
    + 怎样获取到阻塞概要信息？（SetBlockProfileRate）
    + runtime/pprof.Lookup 函数的正确调用方式是什么？
        - runtime/pprof 6 个预定义概要名称: goroutine、heap、allocs、threadcreate、block、mutex
    <img src="img/pprof.png"> 
    + 如何为基于 HTTP 协议的网络服务添加性能分析接口？（net/http/pprof）
    + runtime/trace 代码包的功用是什么？


## 为什么 go 语言没有继承

        面向对象编程，至少在最知名的语言中，涉及到太多关于类型之间关系的讨论，这些关系通常可以自动派生。而 Go 则采用了不同的方法。
        在 go 中，并不要求程序员提前声明两个类型是相关的，go 会自动满足指定其方法子集的任何接口。但这种方法真正的优势可不仅是减少记录。go 的类型可以同时满足多个接口，而不需要传统的多重继承的复杂性。接口可以是非常轻量的，具有一个甚至零个方法的接口就可以表示一个有用的概念。如果出现了新的需求或者用于测试，可以在实体之后直接添加接口，而不需要注释原类型。由于类型和接口之间没有明确的关系，所以也不需要管理或争论的类型层次结构。
        这些思想可以用于构建类似于类型安全的Unix管道的东西。例如，可以参考 fmt.fprintf 是如何格式化打印到任何输出而不仅是文件、bufio 包如何与文件 I/O 完全分离、image 包如何生成压缩文件。所有这些想法都源于用单个接口（io.writer）表示单个方法（write）。这些还只表面上，Go的接口对程序的结构还有着很深远的影响。
        熟悉这些需要一些练习习惯，但这种类型依赖的隐式风格是 Go 高效的事情之一。

> Why is there no type inheritance? https://golang.org/doc/faq#inheritance

# 内建函数
> close()
用于关闭 channel（双向或者只写的），将不在阻塞读取。如果 channel 关闭且没有值，读取 ok 将会为 false
> select case
1. 存在 default 不阻塞
2. 不存在 default 阻塞

# TCMalloc

# 垃圾回收
+ STW (Stop the World): 在标记终止阶段, 会有一个短暂的 STW 阶段, 在这个阶段, 任何剩余的灰色对象会被完成标记, 确保所有存活的对象被标记为黑色.
+ Mark STW，SWEEP 并行
+ 三色标记法
    - 白色对象: 未被垃圾回收器访问过的对象，可能死亡的对象。在 GC 开始时，所有对象都被初始化为白色。
    - 灰色对象: 已被垃圾回收器访问过的对象, 但仍有未被扫描的指针指向白色对象。
    - 黑色对象: 已被垃圾回收器访问过的对象, 且所有字段都已被扫描, 不存在指向白色对象的指针。
+ 写屏障( write barrier): 一种内存保护机制, 用于在修改对象的引用关系时, 保证垃圾回收器能及时捕捉到变化. 在每次引用更新时候增加一段代码, 在引用变更时通知垃圾回收器, 使其更新标记信息, 如在三色标记法中发生对象引用修改操作, 对象处于白色状态, 则写屏障将对象标记为灰色, 确保后续被垃圾回收器正确处理.
    - 前写屏障: 对象引用被更改前执行, 记录引用变更前的信息. 一般用于保持旧的引用关系, 防止重要对象被过早回收.
    - 后写屏障: 对象引用被修改之后执行, 确保引用对象被正确跟踪. 常用于三色标记法, 确保新创建或修改后的对象及时被标记为可达状态.
+ 并发标记: 它允许垃圾收集器在程序继续执行的同时识别哪些内存仍在使用、哪些可以被回收。并发标记通过三色标记法实现，借助写屏障保证准确性。
+ 分代收集: 是一种基于对象生命周期特性的垃圾回收策略. 它引用了大多数对象在创建后不久变得不可达这一观察结果. 根据这个特性, 内存中的对象被划分为不同的代, 对不同代采取不同的垃圾回收策略. 如 Minor GC(年轻代收集)、Major GC/Full GC(老年代收集/全堆收集)、对象晋升。


> 三色标记法
1. 所有对象最开始是白色
2. 从 root 开始找到所有可达对象，标记为灰色，放入待处理队列
3. 遍历灰色对象队列，将其引用对象标记为灰色放入待处理队列，自身标记为黑色
4. 处理完灰色对象队列，执行清扫工作。
5. 完成所有对象扫描后，白色对象即为不可达的"垃圾"，将被回收期清除。

# struct tag
可以很方便的进行 json 、yaml 等文件的解析

# 语言设计
1. 类型设计的原则

        变量包括 (type,value) 两部分,type 包括 static type 和 concrete type.
        reflect/type.go 中的 rtype 与 runtime/type.go 中的 _type 保持一致.

2. golang 反射为什么会慢？

        涉及内存分配以及后续的 GC
        reflect 实现里面有大量的枚举，也就是 for 循环，比如类型之类的

3. goroutine 与线程的协同机制有哪些？

    <img src="img/routine_sync.jpg">
    各个操作系统的系统机制
    <img src="img/os_sync.png">


# golang 问题
1. map、struct、切片在 64 位机器中占用多少字节？
2. 为什么 nil != nil ？
3. 进程、线程、协程有什么关系？
    <img src="../img/routine_thread_process.jpg">

4. 什么是 interface？
    - interface 是具有一组方法的类型，如果一个类型实现了一个 interface 的所有方法，就说该类型实现了 interface. 

5. 原子操作有什么好处？
    - 原子操作可以用互斥体完成，但它比互斥体更快，它是 CPU 而非操作系统提供的能力，如

            var val int32
            ...
            newval = atomic.AddInt32(&val, delta)

6. 用锁有什么需要注意的？
   - 锁不是很容易控制，忘记 unlock 将会导致灾难性后果. 锁粒度不宜过大，不要在锁里面执行费时操作. 读操作阻止写而不阻止读，写操作阻止一切.
7. 使用 golang map 有什么需要注意的?
8. mutex 和 rwmutex 有什么区别?
9. WaitGroup 的作用与原理?
10. 什么是 Context 包?

#  工具
1. go-callvis 源码分析
2. <a href="https://maiyang.me/post/2018-09-14-tips-vscode/">VS Code 中的代码自动补全和自动导入包</a>
    - gocode (auto-completion)
    - gopkgs (auto-completion of unimported packages && Add Import feature)
    - go-outline (Go to symbol in file)
    - go-symbols (Go to symbol in workspace)
    - guru (Find all references and Go to implementation of symbols)
    - gorename (Rename symbols)
    - dlv (Debugging)
    - gocode-gomod (Autocompletion, works with Modules)
    - godef (Go to definition)
    - golint (Linter)
    - gopls (Language Server from Google)
    - gotests (Generate unit tests)
    - gomodifytags (Modify tags on structs)
    - impl (Stubs for interfaces)
    - fillstruct (Fill structs with defaults)
    - goplay (The Go playground)
    - godoctor (Extract to functions and variables)

# ref
1. <a href="http://legendtkl.com/2017/04/28/golang-gc/">Golang 垃圾回收剖析</a>
2. <a href="https://www.jianshu.com/p/c4ec92afeca8">golang 自定义 struct 字段标签</a>
3. <a href="https://www.kancloud.cn/kancloud/effective/72199">Effective Go 中文版</a>
