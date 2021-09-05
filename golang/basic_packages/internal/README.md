# 简介
internal 是 golang 源码的内部 package, . golang 中以 internal 命名的文件夹内部的 API 不可以被外部使用.

## 源码
总计 18000 多行代码，除去测试代码 13000 多行.
+ internal/bytealg   
    - 字节切片相关的函数、如比较、计数等      
+ internal/cpu   
    - 实现了对处理器特性的检测
+ internal/fmtsort   
    - 提供对 map 的稳定排序机制
    - 通过 reflect 获取类型及值
    - 通过 sort.Stable() 进行排序 
+ internal/goroot   
    - func IsStandardPackage(goroot, compiler, path string) bool
    - 判断路径是否为标准库        
+ internal/goversion  
    - golang 版本     
+ internal/lazyregexp  
    - 对 regexp 的封装
+ internal/lazytemplate    
    - 对 text/template.Template 的封装
+ internal/nettrace       
    - 提供了内部的 hooks ，用于追踪 net 包的活动
    - 提供给 net/http/httptrace 使用
+ internal/oserror     
    - 定义给 os 包使用的 errors
+ internal/poll   
    - 支持在有 polling 的文件描述符上的非阻塞 I/O 
    - goroutine 级别阻塞
    - 供 net 和 os 包使用
+ internal/race   
    -  提供帮助函数，用于代码的 race 检测器
+ internal/reflectlite    
    - 提供轻量级 reflect
    - 它的实现仅使用 unsafe 和 runtime 两个标准库
+ internal/singleflight    
    - 提供了用于抑制函数重复调用的机制
+ internal/syscall   
    - unix / windows
+ internal/testenv       
    - 为 golang 团队提供一些测试信息，关于函数是否可用的信息
+ internal/testlog        
    - 在 tests 和 os 包之间，提供 back-channel 交互路径
    - 供 cmd/go 查看环境变量和文件
+ internal/trace
    - 用于追踪的一些接口、方法
+ internal/xcoff
    - 实现了用于访问 XCOFF (Extended Common Object File Format) 类型的文件


## 应用
仅供 golang 标准库使用

## ref
1. [Go 1.4 “Internal” Packages](https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit)