# 反射
1. 基础库简介
2. 优秀框架
3. 源码
4. 应用
6. struct 及 interface
7. gin.Bind 中的反射
8. gorm 中的反射

## reflect 优秀代码
- gorm
- json
- yaml
- gRPC

## golang 类型设计
+ 变量（type,value）
+ interface 的 pair (value,type)
    - reflect.TypeOf 获取 type
    - reflect.ValueOf 获取 value

            type T struct {
                A int    `json:"aaa" test:"testaaa"`
                B string `json:"bbb" test:"testbbb"`
            }

## reflect 核心
+ TypeOf(i interface{}) Type
    - Align() int
    - FieldAlign() int
    - Method(int) Method
    - MethodByName(string)(Method, bool)
    - NumMethod() int
    - Name() string
    - PkgPath() string
    - Size() uintptr
    - String() string
    - Kind() Kind
    - Implements(u Type) bool
    - AssignableTo(u Type) bool
    - ConvertibleTo(u Type) bool
    - Comparable() bool
    - Bits() int
    - ChanDir() ChanDir
    - IsVariadic() bool
    - Elem() Type
    - Field(i int) StructField
    - FieldByIndex(index []int) StructField
    - FieldByName(name string) (StructField, bool)
    - FieldByNameFunc(match func(string) bool) (StructField, bool)
    - In(i int) Type
    - Key() Type
    - Len() int
    - NumField() int
    - NumIn() int
    - NumOut() int
    - Out(i int) Type

+ ValueOf(i interface{}) Value
    - 部分方法如下:
    - Call(in []value) []Value
    - Elem() Value
    - Field(i int) Value
    - FieldByName(name string) Value
    - Index(i int) Value
    - IsNil() bool
    - IsValid() bool
    - MapIndex(key Value)
    - MapKeys() []Value
    - NumFiled() int
    - Set(x Value)
    - Type() Type


## reflect 性能
+ 涉及内存分配以后的 GC
+ reflect 实现中存在 for 循环
