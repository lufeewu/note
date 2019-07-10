# 简介
标准库 unsafe 提供了一些跳过 go 语言类型安全限制的操作

## 源码
仅一个 unsafe.go 文件，总计 206 行，大部分是注释
+ type ArbitraryType
+ type Pointer
+ func Sizeof(v ArbitraryType) uintptr
+ func Alignof(v ArbitraryType) uintptr
+ func Offsetof(v ArbitraryType) uintptr

unsafe 函数的调用始终在编译时求值. 它们返回的结果可以分配给常量.
golang 出于安全考虑不允许以下类型直接转换，但通过 unsafe 这是可能的:
+ 两个不同指针类型的值，例如 int64和 float64.
+ 指针类型和uintptr的值.

unsafe 规则:
+ 任何类型的指针值都可以转换为 unsafe.Pointer.
+ unsafe.Pointer 可以转换为任何类型的指针值.
+ uintptr 可以转换为 *unsafe.Pointer.
+ unsafe.Pointer 可以转换为uintptr.

## unsafe.Pointer 与 unitptr 
+ uintptr是一个整数类型
    - 即使uintptr变量仍然有效，由uintptr变量表示的地址处的数据也可能被GC回收.
+ unsafe.Pointer是一个指针类型
    - unsafe.Pointer值不能被取消引用
    - 如果unsafe.Pointer变量仍然有效，则由unsafe.Pointer变量表示的地址处的数据不会被GC回收
    - unsafe.Pointer是一个通用的指针类型，就像* int等


## 应用
1. cgo 使用时经常用到
2. 类型强制转换
2. sync/atomic 中指针相关的函数

## ref
1. [[译]Go里面的unsafe包详解](https://gocn.vip/question/371)
2. []()