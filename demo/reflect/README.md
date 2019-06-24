# reflect
reflect 是 golang 的一个标准库, 通过反射可以获取变量的类型、值、tag 等，它是实现 gorm、json、yaml 等库的基础. golang 语言通过反射可以在运行时动态的调用对象的方法和属性，可以用来检测存储在变量内部（值 value，类型 concrete type）pair（value, type) 对的机制，官方自带的 reflect 包就是反射包. 


## golang struct
相比与其它语言，golang 的 struct 中可以在类型后增加一段字符串，即 Tag，在编写程序时，常常通过如下形式对数据进行解析等处理，比如下图所示:

        type jsonStruct struct {
            A int `json:"a,omitempty"`
        }

通过 json.Unmarshal 可以将 json 字符串解析到 jsonStruct 类型对象中，同样可以将该类型对象格式化为 json 字符串.

+ field tag 成员标签变量
    - 标签键值对 key:"value"

+ ValueOf
    - Interface()、CanInterface()、Set() ... 
+ TypeOf
    - 通过 TypeOf 可以获取 struct 的结构信息，包括名称、类型、附加 tag
    - StructField
        - Name、Type、Tag
        - Index、Anonymous、PkgPath、Offset
    - Tag/StructTag
        - LookUp
        - Get

## 优质项目
+ gorm
+ json
+ yaml
+ grpc


# ref
1. https://juejin.im/post/5a75a4fb5188257a82110544