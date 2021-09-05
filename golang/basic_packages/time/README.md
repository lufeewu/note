# 简介
time 是 golang 标准库之一，它是十分常用的标准库. 提供了时间显示、测量的函数，提供了定时器等. 日历采用的是公历.

## 源码
总计 9300 多行代码, 除去测试代码 5200 多行.
+ type Weekday int
+ type Month int
+ type Location struct
+ type Time struct 
    - func (t Time) Year() int
    - func (t Time) After(u Time) bool
    - func (t Time) Local() Time
    - func (t Time) UTC() Time
    - func (t Time) Date() (year int, month Month, day int)
    - func (t Time) Clock() (hour, min, sec int)
    - func (t Time) Weekday() Weekday
    - func (t Time) AddDate(years int, months int, days int) Time
    - func (t Time) Sub(u Time) Duration
    - func (t Time) Round(d Duration) Time
    - func (t Time) Format(layout string) string
    - ...
+ type Duration int64
    - Duration 为两个时间点之间经过的时间，以纳秒为单位
    - func Since(t Time) Duration  
    - func (d Duration) Hours() float64
    - func (d Duration) String() string
    - ...
+ type Timer struct
    - timer 代表单次时间事件
    - func NewTimer(d Duration) *Timer
    - func AfterFunc(d Duration, f func()) *Timer
    - func (t *Timer) Stop() bool
    - 
+ type Ticker struct
    - Ticker保管一个通道，并每隔一段时间向其传递"tick"
    - func NewTicker(d Duration) *Ticker
    - func (t *Ticker) Stop()
    - func Sleep(d Duration)
    - func After(d Duration) <-chan Time
    - func Tick(d Duration) <-chan Time

Timer 内部调用了  startTimer 、stopTimer , 它们的实现在 runtime/time.go 中
+ func startTimer(*runtimeTimer)
+ func stopTimer(*runtimeTimer) bool

## 应用
1. 定时器 timer、ticker
2. 日期处理


## ref
1. [go中的定时器timer](http://guidao.github.io/go_timer.html)
2. [What does a function without body mean?](https://stackoverflow.com/questions/14938960/what-does-a-function-without-body-mean)
3. [Function declarations](https://golang.org/ref/spec#Function_declarations)