# gomod
gomod 是 golang 1.11 增加的新特性.  Modules 官方定义为

    模块是相关Go包的集合. modules是源代码交换和版本控制的单元. go命令直接支持使用modules, 包括记录和解析对其他模块的依赖性. modules替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件.

## go111module 配置
GO111MODULE 有三个值: off、on 和 auto . 


## go mod 命令的使用
go mod \<command> [arguments]

    download    download modules to local cache
    edit        edit go.mod from tools or scripts
    graph       print module requirement graph
    init        initialize new module in current directory
    tidy        add missing and remove unused modules 
    vendor      make vendored copy of dependencies
    verify      verify dependencies have expected content
    why         explain why packages or modules are needed

+ module
    - 指定包的名字
+ require
    - 指定包
+ replace
+ exclude

# ref
1. go mod 使用 https://juejin.im/post/5c8e503a6fb9a070d878184a
 