# 简介
flag 标准库是 golang 提供的一个命令行参数的解析的包. 


## 源代码
总代码数 1700 多行，除去测试代码 1000 多行，均在 flag.go 代码中
常用结构、函数
+ type Value
+ type ErrorHandling
+ type Flag
- type FlagSet 它是一组 flags 的集合
    - parseOne() (bool, error)  解析一个命令行交互
    - Parse(arguments []string) error  调用 parseOne() 解析所有命令行
+ func Parse()
+ func Parsed() bool
+ func Args() []string
+ func Int(name string, value int, usage string) *int
+ func Int64(name string, value int64, usage string) *int64
+ func String(name string, value string, usage string) *string
+ func Duration(name string, value time.Duration, usage string) *time.Duration
+ ...

## 第三方库
1. github.com/spf13/cobra