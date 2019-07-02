# 简介
math 标准库包含了一些基本的数学常数和数学函数.

## 源码结构
math 标准库总计 go 代码 34000 多行，除去测试代码 16000 多行. 主要代码是 math/big 处理大数字的多精度计算，总计超过 18000 行代码，不计测试代码约 8200 多行. math/bits、math/cmplx、math/rand 三个子库的代码均约 2000 行. 剩余则涉及是公用函数如 max、min、erf、sqrt、exp、floor、mod、abs、ceil、log 等数学运算函数，但它们主要是对 float64 进行操作, 以及一些常数如自然常数 e、圆周率 pi 等.

## float64 类型
go 语言提供的函数主要都是针对 float64 类型的，如 max、min 等，这个理由主要包括:
1. 由于float64类型要处理infinity和not-a-number这种值，而他们的处理非常复杂，一般用户没有能力，所有go需要为用户提供系统级别的解决办法。
2. 对于int/int64类型的数据，min/max的实现非常简单直接，用户完全可以自己实现. 

也就是说，go 的标准库 math 的操作是为用户提供了较复杂的函数，简单的处理交由用户自己完成.

## 模块
- big
    - type Int 
    - type Rat 代表一个任意精度有理数
    - 实现了大数字的多精度计算
- bits
    - 实现了位的计数和操作函数，如 Reverse、RotateLeft、Len、Div 等
- cmplx
    - cmplx 提供了复数的常用常数和常用函数.
    - 针对 complex128
- rand
    - 实现了伪随机数生成器. 随机数从资源生成（默认公共资源，在程序运行时产生固定序列）.
    - 可通过 seed 函数初始化使得每次运行产生不同序列.
    - type Source
    - type Rand
        - 支持 int 、Uint、Float 等类型
        - 支持正态分布 NormFloat64、指数分布 ExpFloat64
        - 支持伪随机排列切片
    - type Zipf
        - 服从奇普夫分布的随机数
    - func Seed、Int、Uint、Float、NormFloat64、ExpFloat、Perm


## ref
1. [go语言为什么没有min/max(int, int)函数](https://studygolang.com/articles/11545)
2. [Don't abuse math.Max / math.Min](https://mrekucci.blogspot.com/2015/07/dont-abuse-mathmax-mathmin.html)