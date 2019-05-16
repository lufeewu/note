# reflect
reflect 是 golang 的一个标准库, 通过反射可以获取变量的类型、值、tag 等，它是实现 gorm、json、yaml 等库的基础. golang 语言通过反射可以在运行时动态的调用对象的方法和属性，可以用来检测存储在变量内部（值 value，类型 concrete type）pair（value, type) 对的机制，官方自带的 reflect 包就是反射包. 


## 结构
+ ValueOf
+ TypeOf

## 优质项目
+ gorm
+ json
+ yaml
+ grpc




# ref
1. https://juejin.im/post/5a75a4fb5188257a82110544