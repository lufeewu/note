# 简介
golang 的原生自带了关于单元测试的许多工具包。单元测试可以极大的提升代码质量，并能提升整体研发效率。

## mock 库
在编写单元测试中，一些依赖函数需要进行 Mock 处理。golang 官方提供了 [gomock 工具](https://github.com/golang/mock) 工具帮助函数的 mock，这个 mock 方法主要是对 interface 进行 mock。gomock 工具提供了 mockgen 工具自动生成 mock 代码。

对一个 helloword 工程，通过 proto 生成 pb 文件后，可以通过 mockgen 指定 interface 生成对应 pb 的 mock 文件。命令示例如下

        mockgen -destination greeter_mock.go -package helloworld .  GreeterClientProxy


# 参考
1. [https://github.com/golang/mock](https://github.com/golang/mock)