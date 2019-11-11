## golang 基础
1. 值类型与引用类型
2. go 语言的左值、右值(表示存储在存储器某个地址的数据值)
3. 内置函数 new、make、copy、delete、close、len、cap、new、make、copy、append
4. 


## golang 进阶
1. golang 的可寻址(addressable)
2. golang 的调度器 GPM
3. golang 的 CSP 并发模型
4. golang 的值类

## Uber golang 规范
- 通用准则
1. [Effective Go](src="https://golang.org/doc/effective_go.html")
2. [The Go common mistakes guide](src="https://github.com/golang/go/wiki/CodeReviewComments")
3. golint 、 go vet 检查源码
4. goimports

- 指导性原则
1. 指向 interface 的指针
2. 值接收器、指针、值
3. 零值 Mutex、指针、结构体嵌入
4. 引用类型 slice、Map 的拷贝
5. 使用 defer 做清理
6. channel 的 size 确定
7. iota 枚举初始值
8. 错误类型(errors.New、fmt.Errorf、wrapped error、”pkg/errors”.Wrap)
9. Error Wrapping 
10. 处理类型断言 "comma ok"
11. 生产环境中避免 panic

- 性能
1. strconv 比 fmt 快
2. 避免字符串到字节的反复转换

- 样式
1. 相似声明放在一组
2. import 组内导包顺序
3. 包命名规则（全小写、无下划线、简短简洁、无复数、信息量足)
4. 函数命名约定 [MixedCaps](src="https://golang.org/doc/effective_go.html#mixed-caps")
5. 包导入别名（除冲突外避免别名）
6. 函数分组与顺序
7. 减少嵌套
8. 减少不必要的 else 
9. 顶层变量申明
10. 对未导出的顶层常量、变量，使用 _ 作为前缀
11. 结构体中的嵌入
12. 使用字段名初始化结构体
13. 本地变量申明（短变量申明式 :=，空切片 slice）
14. nil 是一个有效的长度为 0 的 slice
15. 缩小变量作用域
16. 避免裸参数
17. 使用原始字符串字面值``并避免转义
18. 初始化结构体引用使用 &T{} 代替 new(T)
19. 格式化字符串放在 Printf 外部

- 模式
1. 测试表，表格驱动测试与子测试一起使用
2. 功能选项（一种模式）


## golang 文章记录
1. [Frequently Asked Questions (FAQ)](src="https://golang.org/doc/faq")
2. [Effective Go](src="https://golang.org/doc/effective_go.html")
3. [Uber Go语言编码规范](src="https://tonybai.com/2019/10/12/uber-go-style-guide/")
4. [Uber Go语言编码规范 github](src="https://github.com/uber-go/guide/blob/master/style.md")
5. [Go Code Review Comments](src="https://github.com/golang/go/wiki/CodeReviewComments")