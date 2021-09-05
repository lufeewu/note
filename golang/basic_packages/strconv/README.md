# 简介
strconv 作为 golang 的标准库实现了基本的数据类型和其字符串表示的相互转换.

## 源码
总计 7400 多行，除去测试代码 4300 多行
+ func Atoi(s string) (i int, err error)
+ func Itoa(i int) string
+ func ParseInt(s string, base int, bitSize int) (i int64, err error)
+ func Quote(s string) string
    - 返回字符串在 go 语言下的双引号字面值表示，控制字符、不可打印字符会进行转义. 如 \t、\n、\xFF、\u0100 等
+ func IsPrint(r rune) bool
+ func AppendBool(dst []byte, b bool) []byte


## string 与 []byte
在 builtin 包中，可以看到 string 是是一个 8-bit bytes 的字符集合，string 可以是空的，但空不代表是 nil. 它们之间可以相互转换，且数据不会变化. 它们的区别是 string 是不可变的，但是 byte 切片是可变的.

## ref
1. [浅谈 Go 语言实现原理](https://draveness.me/golang/datastructure/golang-string.html)