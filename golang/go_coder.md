# go 语言进阶之路
涵盖 go 语言的基础入门、web 编程、原理剖析、优化进阶以及 golang 学习的资料.

## 目录
+ 第一章 环境准备
    - GOPATH 配置
    - Go 命令
    - 编写 HelloWorld 程序 
        - main 包、main 函数
    - 开发工具 vscode

+ 基础知识（语言的使用）
    - 基础知识点
        - 25 个关键字、37 个预定义标识符、41 个标准库
    - 基本运算符
        - 算术运算符、关系运算符、逻辑运算符、位运算符、赋值运算符、其它运算符
    - 条件与循环
        - if、else、else if
        - for、range
    - package 管理（mod、dep）
    - 类型系统
        - 结构体与函数
        - 接口 interface 使用
        - 类型转换
        - 数组、slice 与 map
    - 内置函数
    - 并发三件套 go、select、channel
        - 实战 - 生产者与消费者
    - 基础面试题
        - defer
        - slice
    - 编码规范
        - lint、vet 
        - 边界检查 bce
+ Web 工程（web 的开发）
    - 常用框架
        - 原生 http 库
        - gin 框架
        - gorm 框架
    - 实战 - 搭建一个简单的 web 服务
+ 原理剖析
    - go runtime 解析
        - 内存分配
        - gpm 模型
        - 垃圾回收
    - golang 引用类型的底层实现
        - slice
        - map
        - channel
+ 高阶用法
    - pprof
    - race
    - testing
+ go 生态
    - go proposal
+ 学习推荐
    - go web 编程
    - go 语言原本[欧长坤, 欧神]
    - go 夜读（tidb 杨文）