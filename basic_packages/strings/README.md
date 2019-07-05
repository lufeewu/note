# 简介
strings 作为 golang 标准库之一，包提供了用于操作字符串的简单函数.

## 源码
总计代码 7400 多行，除去测试代码 4300 多行.
+ type Reader
    - 实现了 io.Reader、io.ReaderAt,、io.Seeker、io.WriterTo、io.ByteScanner 及 io.RuneScanner
    - 从一个字符串里读取
    - NewReader(s string) *Reader { return &Reader{s, 0, -1} }
+ type Replacer
    - 用于一系列字符串的替换
    - 多个 goroutine 并发安全
+ func HasPrefix(s, prefix string) bool
+ func HasSuffix(s, suffix string) bool
+ func Contains(s, substr string) bool
+ func Index(s, sep string) int
+ func IndexFunc(s string, f func(rune) bool) int
+ func LastIndex(s, sep string) int
+ func ToLower(s string) string
+ func ToUpper(s string) string
+ func Repeat(s string, count int) string
+ func Replace(s, old, new string, n int) string
+ func Map(mapping func(rune) rune, s string) string
    - 对字符串 s 进行映射
+ func TrimSpace(s string) string
+ func TrimRight(s string, cutset string) string
+ func Split(s, sep string) []string
+ func Join(a []string, sep string) string
+ ...


