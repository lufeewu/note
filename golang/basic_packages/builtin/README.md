# 简介
builtin 作为 golang 的基础库之一. 主要的意义是提供 golang 预定义标识符提供文档.

## 源码
仅一个 builtin.go 文件，262 行代码
+ type bool
+ type byte
+ type rune
+ type int
+ type int8
+ type int16
+ type int32
+ type int64
+ type uint
+ type uint8
+ type uint16
+ type uint32
+ type uint64
+ type float32
+ type float64
+ type complex64
+ type complex128
+ type uintptr
+ type string
+ type error
+ type Type
+ type Type1
+ type IntegerType
+ type FloatType
+ type ComplexType
+ func real(c ComplexType) FloatType
+ func imag(c ComplexType) FloatType
+ func complex(r, i FloatType) ComplexType
+ func new(Type) *Type
+ func make(Type, size IntegerType) Type
+ func cap(v Type) int
+ func len(v Type) int
+ func append(slice []Type, elems ...Type) []Type
+ func copy(dst, src []Type) int
+ func delete(m map[Type]Type1, key Type)
+ func close(c chan<- Type)
+ func panic(v interface{})
+ func recover() interface{}
+ func print(args ...Type)
+ func println(args ...Type)