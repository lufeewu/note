# 简介
标准库 regexp 实现了正则表达式搜索，使用 RE2 语法，和 Perl、Python 等语言的正则基本保持一致.


## 源码
总计 9600 多行代码，除去测试代码 3600 多行.
+ type Regexp struct
    - 代表一个编译好的正则表达式
    - 多线程安全
    - func (re *Regexp) String() string
    - func (re *Regexp) Match(b []byte) bool
    - func (re *Regexp) MatchString(s string) bool
    - func (re *Regexp) MatchReader(r io.RuneReader) bool
    - func (re *Regexp) FindAll(b []byte, n int) [][]byte
    - func (re *Regexp) FindAllString(s string, n int) []string
    - func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
    - func (re *Regexp) ReplaceAllString(src, repl string) string
    - func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
    - func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
    - ...
+ func Match(pattern string, b []byte) (matched bool, err error)
+ func MatchString(pattern string, s string) (matched bool, err error)
+ func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
+ regexp/syntax 
    - syntax 包将正则表达式解析为解析树，并编译成程序
    - func Parse(s string, flags Flags) (*Regexp, error)
    - type Regexp struct
        - func (re *Regexp) CapNames() []string
        - func (x *Regexp) Equal(y *Regexp) bool
        - func (re *Regexp) MaxCap() int
        - func (re *Regexp) Simplify() *Regexp
        - func (re *Regexp) String() string

## 应用
正则表达式在文本处理中是十分常用的.
1. Gin 的 router 模块
2. Gorm 的

## ref
1. [正则 re2 语法](http://code.google.com/p/re2/wiki/Syntax)