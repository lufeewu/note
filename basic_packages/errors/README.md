# 简介
errors 基础库是 golang 提供的错误处理库. 它实现了基本的创建错误值函数.

## 源码
errors 基础库包含 440 多行代码. 除去测试代码仅 106 行.
+ errors.go
    - 实现了 error 接口的 errorString 类
    - type errorString
    - func (e *errorString) Error() string 
    - func New(text string) error
+ wrap.go
    - func Unwrap(err error) error 返回 Unwrap 方法的实现结果
    - func Is(err, target error) bool
    - func As(err error, target interface{}) bool 
    - interface { Unwrap() error }
    - interface{ Is(error) bool }
    - interface{ As(interface{}) bool }

## 第三方 errors 库
1. github.com/pkg/errors