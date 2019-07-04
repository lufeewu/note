# 简介
expvar 作为标准库之一，提供了公共变量的标准接口，如服务的操作计数器. 

## 源码
总计 960 多行代码，除去测试代码 360 多行，只有 expvar.go 文件
+ type Var
+ type Int
+ type Float
+ type String
+ type Func
+ type KeyValue
+ type Map
+ Get(name string)Var
+ Publish(name string, v Var)
+ Do(f func(KeyValue))

expvar 包含默认的变量用于记录所有 published 的变量
+ All published variables.

        var (
            vars      sync.Map // map[string]Var
            varKeysMu sync.RWMutex
            varKeys   []string // sorted
        )

## 使用示例
每个请求 handler() 后，在向访问者发送响应消息之前增加计数器. expvar 默认注册了 /debug/vars 的接口以 json 格式导出变量.

    package main

    import (
        "expvar"
        "fmt"
        "net/http"
    )

    var visits = expvar.NewInt("visits")
    func handler(w http.ResponseWriter, r *http.Request) {
        visits.Add(1)
        fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    }

    func main() {
        http.HandleFunc("/", handler)
        http.ListenAndServe(":1818", nil)
    }

使用 expvar 创建的对这些公共变量的读写操作都是原子级的. 另外，expvar 包中默认注册了如下变量, 在调用 /debug/vars 时一并以 JSON 格式返回:
+ cmdline   os.Args
+ memstats  runtime.Memstats