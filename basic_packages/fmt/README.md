# 简介
fmt 标准库是 golang 最常用的库之一. 它实现了格式化的 I/O ，类似于 C 语言 printf 和 scanf.

## 源码
总计 7100 多行代码，除去测试代码仅 3300 多行.

常用函数、结构
+ func Printf(format string, a ...interface{}) (n int, err error)
+ func Scanf(format string, a ...interface{}) (n int, err error)
+ func Sprintf(format string, a ...interface{}) string
+ type Scanner
+ type Stringer

内部核心函数、结构
+ type pp struct
    - doPrintf(format string, a []interface{})
    - printValue(value reflect.Value, verb rune, depth int)
    - printArg(arg interface{}, verb rune)
    - handleMethods(verb rune) (handled bool) 
    - fmt0x64(v uint64, leading0x bool)
    - fmtInteger(v uint64, isSigned bool, verb rune) 
    - fmtFloat(v float64, size int, verb rune)
    - fmtComplex(v complex128, size int, verb rune)
    - fmtString(v string, verb rune)
    - fmtBytes(v []byte, verb rune, typeString string)
    - fmtPointer(value reflect.Value, verb rune)
    - ...


## Printing 和 Scanning
通过 %+char 或 '+'、' '、'-'、'#'、'0' 等的组合，可以输出不同格式的数据，如布尔、整数(不同进制)、浮点数、复数、字符串、[]byte、指针等.
Scan、Scanf 和 Scanln 从标准输入 os.Stdin 读取文本; Fscan、Fscanf、Fscanln 从指定 io.Reader 接口读取文本

## 其它类似
1. log 等输入、输出格式化