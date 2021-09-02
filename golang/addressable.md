# 简介
go 语言规范中规定了可寻址(addressable)对象的定义。

## 可寻址
golang 中可寻址包括以下
- 一个变量&x
- 指针引用 &*x
- slice 的所以操作 &s[1]
- 可寻址的 struct 字段 &point.X
- 可寻址数组的索引操作 &a[0]
- composite literal 类型  &struct{ X int }{1}

## 不可寻址
golang 中不可寻址包括以下
- 字符串中的字节
- map 对象中的元素
- 接口对象的动态值(通过 type assertions 获取)
- 常数
- literal 值
- package 级别的函数
- 方法 method (用作函数值)
- 中间值(intermediate value): 函数调用、显示类型转化、各类型的操作(如通道接受、子字符串操作、子切片操作、加减乘除等运算)



## 参考
1. [Pointers vs. Values](src="https://golang.org/doc/effective_go.html#pointers_vs_values")
2. [go addressable 详解](src="https://colobu.com/2018/02/27/go-addressable/")
3. [Which values can and which values can't be taken addresses?](src="https://go101.org/article/unofficial-faq.html#unaddressable-values")