# 简介
syscall 是 golang 标准库中提供操作系统原语的接口. 它首先会被其它的标准库引用，比如 os、net、time，可以提供可移植的系统接口.

## 源码
总计包含 148000 多行代码，除去测试代码 145000 多行代码，涉及不同平台 windows、unix、linux. 另外包含 28 个多平台的汇编文件.
比如涉及以下一些函数可直接调用:
+ func Getuid() (uid int)
+ func EpollCreate(size int) (fd int, err error)
+ func Clearenv()
+ func Chmod(path string, mode uint32) (err error)
+ func Chown(path string, uid int, gid int) (err error)
+ func Kill(pid int, sig Signal) (err error)
+ func Mkdir(path string, mode uint32) (err error)
+ func Seek(fd int, offset int64, whence int) (off int64, err error)
+ func Setgid(gid int) (err error)
+ func StringByteSlice(s string) []byte
+ ...
另外也提供了一些系统相关类型、接口
+ type Signal int
+ type TCPInfo struct
+ type Sysinfo_t struct
+ ...

## 应用
1. 对系统的基本操作，如文件操作、设置环境变量之类

## ref
1. [Golang 系统调用 syscall ](https://www.jianshu.com/p/3b6935e3bb50)