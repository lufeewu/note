# golang 反射 reflect
reflect 是 golang 的一个标准库, 通过反射可以获取变量的类型、值、tag 等，它是实现 gorm、json、yaml 等库的基础. golang 语言通过反射可以在运行时动态的调用对象的方法和属性.


## 源码结构
golang reflect 反射的源码位于 golang/go/src/reflect 中，目录下包含以下文件

+ **type.go**
+ **value.go**
+ swapper.go
+ deepequal.go
+ makefunc.go
+ all_test.go
+ asm_s390x.s
+ example_test.go   
+ set_test.go
+ export_test.go
+ tostring_test.go
+ asm_amd64p32.s
+ asm_mips64x.s
+ asm_386.s
+ asm_arm.s
+ asm_mipsx.s
+ asm_wasm.s
+ asm_amd64.s
+ asm_arm64.s
+ asm_ppc64x.s

在源码中, .s 是一些 golang 汇编文件，核心代码 type.go 3155 行代码, value.go 包含 2774 行代码，swapper.go 及 deepequal.go makefunc.go 总计 309 行代码，剩余 go test 源码 7536 余行. 
可以看到 reflect 库的核心源码仅 6000 余行, 对这些代码的理解, 有助于更深刻的理解 golang 的 struct 及一些优秀框架、模块的实现原理.


## golang struct
相比与其它语言, golang 的 struct 中可以在类型后增加一段字符串, 即 Tag，在编写程序时，常常通过如下形式对数据进行解析等处理，比如下图所示:

        type jsonStruct struct {
            A int `json:"a,omitempty"`
        }

上面的 struct 中, 变量包含 value、type 及 tag, 其中 tag 内的成员标签变量时 key、value 健值对形式的，参考 StructTag.Lookup(key string) 中获取 tag 中指定 key 的方法 

通过 json.Unmarshal 可以将 json 字符串解析到 jsonStruct 类型对象中，同样可以将该类型对象格式化为 json 字符串. 而 json.Unmarshal 及 json.Marshal 都是通过 reflect 实现的.

## 基础库 reflect 
反射是用来检测存储在变量内部（值 value; 类型 concrete type) pair 对的一种机制. golang reflect 提供了两种类型用于访问接口变量的内容, reflect.ValueOf() 和 reflect.TypeOf()，分别用于获取 pair 中的 value 和 type. 示例如下

    package main

    import (
        "fmt"
        "reflect"
    )

    func main() {
        var num float64 = 1.2345

        fmt.Println("type: ", reflect.TypeOf(num))
        fmt.Println("value: ", reflect.ValueOf(num))
    }

    运行结果:
    type:  float64
    value:  1.2345


+ ValueOf 可用于获取类型的 value
    - Interface()、CanInterface()、Set() ... 
+ TypeOf 可用于获取类型的 type
    - 通过 TypeOf 可以获取 struct 的结构信息，包括名称、类型、附加 tag
    - StructField
        - Name、Type、Tag
        - Index、Anonymous、PkgPath、Offset
    - Tag/StructTag
        - LookUp
        - Get

## 优质项目
reflect 解决的是运行时获取的对象方法、属性. 在网络中的数据传输，往往时不携带类型信息的字符串、二进制压缩数据. 那么 golang 程序在与其它进程进行通信时，通过 reflect, 便可以很好的完成数据的解析与序列化. 如远程数据库操作、RPC 数据、http 数据、本地 yaml 文件读取等. 相关的这些优质开源库便是通过反射实现，如下面这些 golang 开发中经常用到的框架:
+ gorm
+ json
+ yaml
+ gRPC
+ protobuf
+ gin.Bind()
+ ...

## gin 中反射的应用
对于刚提到的用了 reflect 的优质开源项目，这里以 gin 框架为例，看看是如何将 reflect 运用上去的. 平时我们写代码时是怎样不知不觉的将反射运用到项目中的.

使用 gin 框架时，经常会通过 Bind() 相关的函数将传入的参数转化为结构体. 下面以 gin.Context.Bind() 为例介绍反射是如何被运用到代码中的.

在 gin-gonic/gin/Context.go 的 Bind() -> MustBindWith() -> ShouldBindWith() -> binding.Bind() 找到 Bind() 方法，它是接口(interface) Binding 的一个方法

    type Binding interface {
        Name() string
        Bind(*http.Request, interface{}) error
    }

在 gin-gonic/gin/binding 中,  实现上面 Binding 接口 的 struct 有如下:
+ form/formBinding
+ json/jsonBinding
+ msgpack/msgpackBinding
+ protobuf/protobufBinding
+ query/queryBinding
+ xml/xmlBinding

从名称可以看到，它们就是 http 传输过程中常用的参数传输的格式. 

其中, gin 框架的 form 在实现 Bind() 方法时候，用到了 mapForm 方法，它位于 binding/form_mapping 中的, 而 mapForm 中使用 reflect 实现了 form 数据到 map[string][]string 的转换, 代码如下:

    func mapForm(ptr interface{}, form map[string][]string) error {
        typ := reflect.TypeOf(ptr).Elem()
        val := reflect.ValueOf(ptr).Elem()
        for i := 0; i < typ.NumField(); i++ {
            typeField := typ.Field(i)
            structField := val.Field(i)
            if !structField.CanSet() {
                continue
            }

            structFieldKind := structField.Kind()
            inputFieldName := typeField.Tag.Get("form")
            inputFieldNameList := strings.Split(inputFieldName, ",")
            inputFieldName = inputFieldNameList[0]
            var defaultValue string
            if len(inputFieldNameList) > 1 {
                defaultList := strings.SplitN(inputFieldNameList[1], "=", 2)
                if defaultList[0] == "default" {
                    defaultValue = defaultList[1]
                }
            }
            if inputFieldName == "" {
                inputFieldName = typeField.Name

                // if "form" tag is nil, we inspect if the field is a struct or struct pointer.
                // this would not make sense for JSON parsing but it does for a form
                // since data is flatten
                if structFieldKind == reflect.Ptr {
                    if !structField.Elem().IsValid() {
                        structField.Set(reflect.New(structField.Type().Elem()))
                    }
                    structField = structField.Elem()
                    structFieldKind = structField.Kind()
                }
                if structFieldKind == reflect.Struct {
                    err := mapForm(structField.Addr().Interface(), form)
                    if err != nil {
                        return err
                    }
                    continue
                }
            }
            inputValue, exists := form[inputFieldName]

            if !exists {
                if defaultValue == "" {
                    continue
                }
                inputValue = make([]string, 1)
                inputValue[0] = defaultValue
            }

            numElems := len(inputValue)
            if structFieldKind == reflect.Slice && numElems > 0 {
                sliceOf := structField.Type().Elem().Kind()
                slice := reflect.MakeSlice(structField.Type(), numElems, numElems)
                for i := 0; i < numElems; i++ {
                    if err := setWithProperType(sliceOf, inputValue[i], slice.Index(i)); err != nil {
                        return err
                    }
                }
                val.Field(i).Set(slice)
                continue
            }
            if _, isTime := structField.Interface().(time.Time); isTime {
                if err := setTimeField(inputValue[0], typeField, structField); err != nil {
                    return err
                }
                continue
            }
            if err := setWithProperType(typeField.Type.Kind(), inputValue[0], structField); err != nil {
                return err
            }
        }
        return nil
    }

 对于, json、msgpack、protobuf、xml 等格式的转化，则直接继续调用相关库进行 decode 完成. 具体应用反射的过程，可以进一步查看它们的源码.

## 缺陷
1. 性能问题
    - 涉及内存分配及后续 GC . 每一次的反射运用，都会涉及到 GC. golang GC 是十分低效的，可能会导致不可预料的性能效率问题.
    - reflect 里面有大量枚举, for 循环
2. 代码可读性. 过多的运用反射, 在运行时才通过反射确定类型，可能导致后续维护时困难.

## 总结
在项目中，可能很少会直接使用到 reflect, 但是反射 reflect 作为 golang 40 余个基础库其中之一，使用是及其广泛的，尤其是 golang 的 struct 中经常会用到 Tag 类型，它对应到底层便会使用反射. 另外, 许多优秀的开源程序都会在底层用到 reflect，如上文提到的 gin、gorm、gRPC 等.

理解底层的 reflect, 可以加深对 golang 的结构体 struct 的理解. 在做一些程序设计时，也可在脑海里多提供一种解决问题的思路. 对于各个库之间的关系，也可以更加清晰，有利于梳理技术栈.

开源的框架、组件可能有几百、上千个，但是**基础库只有 40 多个**，它们反而是实现框架、组建、程序等实现的基础. 掌握 golang 的基础库，再阅读优秀的 golang 源码，可以更快速的理解.

# ref
1.  https://juejin.im/post/5a75a4fb5188257a82110544
2.  https://golang.org/pkg/reflect/