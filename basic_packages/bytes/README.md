# 简介
golang 的标准库 bytes 实现了操作 []byte 的常用函数. 实现的函数和 strings 包的函数相当类似。

## 源码
总计代码  5400 多行，除去测试代码仅 1700 多行.
它提供了许多关于 byte 切片的操作.
+ type Reader
    - 实现了 io.Reader、io.Seeker、io.ReaderAt、io.WriterTo、io.ByteScanner、io.RunnScanner 接口.
    - 通过一个 []byte 读取数据.
+ type Buffer
    - 实现了读写方法可变大小的字节缓冲. 
    - 该类型的零值是可以用于读写的缓冲.
+ 公共方法
    - Compare(a, b []byte) int
    - Equal(a, b []byte) bool
    - EqualFold(s, t []byte) bool
    - Runes(s []byte) []rune
    - HasPrefix(s, prefix []byte) bool
    - HasSuffix(s, suffix []byte) bool
    - Contains(b, subslice []byte) bool
    - Count(s, sep []byte) int
    - Index(s, sep []byte) int
    - IndexByte(s []byte, c byte) int
    - IndexRune(s []byte, r rune) int
    - IndexAny(s []byte, chars string) int
    - Index-(s []byte, f -(r rune) bool) int
    - LastIndex(s, sep []byte) int
    - LastIndexAny(s []byte, chars string) int
    - LastIndex-(s []byte, f -(r rune) bool) int
    - Title(s []byte) []byte
    - ToLower(s []byte) []byte
    - ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte
    - ToUpper(s []byte) []byte
    - ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte
    - ToTitle(s []byte) []byte
    - ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte
    - Repeat(b []byte, count int) []byte
    - Replace(s, old, new []byte, n int) []byte
    - Map(mapping -(r rune) rune, s []byte) []byte
    - Trim(s []byte, cutset string) []byte
    - TrimSpace(s []byte) []byte
    - Trim-(s []byte, f -(r rune) bool) []byte
    - TrimLeft(s []byte, cutset string) []byte
    - TrimLeft-(s []byte, f -(r rune) bool) []byte
    - TrimPrefix(s, prefix []byte) []byte
    - TrimRight(s []byte, cutset string) []byte
    - TrimRight-(s []byte, f -(r rune) bool) []byte
    - TrimSuffix(s, suffix []byte) []byte
    - Fields(s []byte) [][]byte
    - Fields-(s []byte, f -(rune) bool) [][]byte
    - Split(s, sep []byte) [][]byte
    - SplitN(s, sep []byte, n int) [][]byte
    - SplitAfter(s, sep []byte) [][]byte
    - SplitAfterN(s, sep []byte, n int) [][]byte
    - Join(s [][]byte, sep []byte) []byte

