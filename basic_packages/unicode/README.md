# 简介
标准库 unicode 是 golang 提供的一些数据及函数，用来测试 Unicode 编码的一些属性. Unicode 即国际码，它对世界上大部分文字系统进行了整理、编码.

## 源码
总计 11000 多行，除去测试代码 8900 多行，其中包含 7700 多行 Unicode 表
+ type CaseRange
    - 代表简单的 unicode 码值的一一映射
+ type Range16
    - 代表一系列16位 unicode 码值
+ type Range32
    - 代表一系列32位 unicode 码值
+ type RangeTable
    - 通过列出集合中码值的范围，定义了一个 unicode 码值的集合
+ type SpecialCase
    - SpecialCase 代表特定语言的字符映射
+ func Is(rangeTab *RangeTable, r rune) bool
+ func In(r rune, ranges ...*RangeTable) bool
+ func IsOneOf(ranges []*RangeTable, r rune) bool
+ func IsSpace(r rune) bool
+ func IsDigit(r rune) bool
+ func IsNumber(r rune) bool
+ func IsLetter(r rune) bool
+ func IsGraphic(r rune) bool
+ func IsControl(r rune) bool
+ func IsMark(r rune) bool
+ func IsPrint(r rune) bool
+ func IsPunct(r rune) bool
+ func IsSymbol(r rune) bool
+ func IsLower(r rune) bool
+ func IsUpper(r rune) bool
+ func IsTitle(r rune) bool
+ func To(_case int, r rune) rune
+ func ToLower(r rune) rune
+ func ToUpper(r rune) rune
+ func ToTitle(r rune) rune
+ func SimpleFold(r rune) rune

golang 中常用 rune 来处理 unicode 和 UTF-8

## 应用

## ref
1. [【golang】浅析rune数据类型](https://juejin.im/post/5b44caebf265da0f491b8b83)