# 简介
log 是 golang 标准库实现的简单日志服务. 另外提供了 syslog ，一个简单的系统日志服务接口. 可以通过它使用 TCP、UDP 或 UNIX 域 socket 发送日志到系统的 syslog 守护进程中.

## 源码
总计 1400 多行代码，除去测试代码 700 多行. syslog 代码总计 790 多行，除去测试代码 370 多行.
+ type Logger  
    - 提供一些格式化输出方法
    - 线程安全
+ var std = New(os.Stderr, "", LstdFlags) 全局 logger
+ func New(out io.Writer, prefix string, flag int) *Logger
    - 创建 logger，并将日志信息写入 io.Writer 中
+ func Flags() int
+ func SetFlags(flag int) 设置输出选项
+ func Prefix() string
+ func SetPrefix(prefix string)
+ func SetOutput(w io.Writer)
+ func Printf(format string, v ...interface{})
+ func Print(v ...interface{})
+ func Println(v ...interface{})
+ func Fatalf(format string, v ...interface{})
+ func Fatal(v ...interface{})
+ func Fatalln(v ...interface{})
+ func Panicf(format string, v ...interface{})
+ func Panic(v ...interface{})
+ func Panicln(v ...interface{})

## 其它日志库
1. github.com/sirupsen/logrus