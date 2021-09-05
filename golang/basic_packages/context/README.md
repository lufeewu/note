# 简介
context 标准库是使用较多的库之一, 它在跨 API 边界和进程之间携带了 deadlines、撤销信号量及其它请求范围值.

## 作用
对于传入服务的请求，应该创建一个 context, 而在接受外部调用的服务方法则应该接受一个 context. 在函数链上则必须传播 context. 另外一可以将它替换为派生的 context 如 WithCancel、WithDeadline、WithTimeout 或者 WithValue. 当一个 context 被 canceled，后面派生的 context 都会 canceled. 

## 源码
代码总数 1400 多行，除去测试代码 520 行,均在 context.go 文件中
+ type CancelFunc
+ type Context
    - func Background() Context
    - func TODO() Context
    - func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    - func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
    - func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
    - func WithValue(parent Context, key, val interface{}) Context
