# 简介
golang 提供了很方便的方式编写并发程序. 通过关键字 go 可以快速的启动一个并发的协程. 结合 channel 、标准库中的 sync.WaitGroup 等可以简单快捷的对协程进行管理. 下文将介绍在 go 中管理并发程序的方式.


## 协程管理
有这样一个需求，需要发起若干个网络请求，并且是不同的网络请求，也即函数、返回值会有不同. 对于多个网络请求，若是串行调用，将会在网络 I/O 上较多时延. 所以最好通过串行进行. 通常的思路如下:
对于这里假设只有俩个函数:

        func customFunc1(ctx *context.Context, args1 uint32, args2 string)
        func customFunc2(ctx *context.Context, args1 uint32, args2 string)

那么使用 goroutine 可以如下进行:

        var err1, err2 error
        var wg sync.WaitGroup
        wg.Add(2)
        go func(){
            defer wg.Done()
            err1 = customFunc1(ctx, args1, args2)
        }()
        go func(){
            defer wg.Done()
            err2 = customFunc1(ctx, args3, args4)
        }()
        wg.Wait()
        if err1 != nil {
            // todo
        }
        if err2 != nil {
            // todo
        }

上面的并发控制主要基于 sync.WaitGroup , 因为预先知道是 2 个函数的并发执行, 所以只需要通过 wg.Wait() 等待 goroutine 处理完成, 然后再分别处理两个函数的返回值 err1 和 err2. 

上述的写法是通常的方式，但是当并发的数量、函数的数量并不是 2 个的时候. 上述的写法就会十分臃肿了. 同时对于不定数量的并发执行, 难以有效的处理返回的错误值. 
对于协程内的值传输，go 里面首先想到的就是 channel,

        var errors = make(chan error, 2)

上述的写法有一个问题，就是当 customFunc 数量是未知的时候，这里 errors 该申请的长度是未知的.

下面的代码通过函数类型、并构造函数类型的数组, 通用化的处理 goroutine 的并行, 并在并发完成后统一处理 errors. 


        var funcs []func() error // 函数类型数组
        funcs = append(funcs, func() error {
            return customFunc1(ctx, args1, args2)
        })
        funcs = append(funcs, func() error {
            return customFunc1(ctx, args3, args4)
        })

        var wg sync.WaitGroup
        var errors = make(chan error, len(funcs)) // errors 通道
        wg.Add(len(funcs))
        for _, f := range funcs {
            go func(execFunc func() int) { 
                defer wg.Done()
                err := execFunc()
                if errCode != 0 {
                    errors <- err
                }
            }(f)
        }
        wg.Wait()
        if len(errors) > 0 {
            // todo
        } 

可以看到，上述的代码将需要并行处理的函数统一用 func() error 形式的函数处理，并构造这样的数组. 利用数组的方式在并发前构造并发所需的初始参数 errors、wg 、并统一处理并发过程中的控制，最后在并发结束前的阻塞以及最后的初始化. 

这样的写法，好处是在后续迭代代码时候，若有新增的并发处理，只需在向数组 funcs 中再添加 func() error 函数元素(将需要处理的对象封装到 func() error 中)，那么关于并发处理、错误处理都可以不需修改代码. 只需要在一处利用 append 进行添加即可. 

上述的代码已经可以看到一个简单的 goroutine 管理器了，如果再进一步抽象处理. 就可以管理任意的并发了. 构造一个如下所示的

    type Funcs struct{
        funcs []func() error
    }
    func (f *Funcs) Add(func() error )
    func (f *Funcs) Run(func() error )

对于这样的方式开源工程 https://github.com/oklog/run 便通过 100 行左右代码实现了 goroutine 的管理. 开源监控软件 prometheus 便是通过 oklog 管理初始化的进程的. 有兴趣的可以进一步研究.