# 简介
标准库 path 实现了对由斜杠分隔的**路径**的相关操作函数. path/filepath 则实现了兼容各操作系统的文件路径的操作函数.

## 源码
总计 5000 行代码，除去测试代码 1800 多行.
+ func IsAbs(path string) bool
+ func Split(path string) (dir, file string)
+ func Join(elem ...string) string
+ func Dir(path string) string
+ func Base(path string) string
+ func Ext(path string) string
+ func Clean(path string) string
+ func Match(pattern, name string) (matched bool, err error)
+ path/filepath
    - 兼容不同操作系统
    - func IsAbs(path string) bool
    - func Abs(path string) (string, error)
    - func Rel(basepath, targpath string) (string, error)
    - func SplitList(path string) []string
    - func Split(path string) (dir, file string)
    - func Join(elem ...string) string
    - func FromSlash(path string) string
    - func ToSlash(path string) string
    - func VolumeName(path string) (v string)
    - func Dir(path string) string
    - func Base(path string) string
    - func Ext(path string) string
    - func Clean(path string) string
    - func EvalSymlinks(path string) (string, error)
    - func Match(pattern, name string) (matched bool, err error)
    - func Glob(pattern string) (matches []string, err error)
    - type WalkFunc func(path string, info os.FileInfo, err error) error
    - func Walk(root string, walkFn WalkFunc) error
    - func HasPrefix(p, prefix string) bool

## 条件编译
+ 编译标签
    - 编译标签可以是多个，它们可以是与或非的逻辑关系
    - 例子 : // +build darwin freebsd netbsd openbsd
+ 文件后缀
    - filepath 中通过 _windows、_unix 的后缀区分不通的系统，使得编译的时候只会使用一种文件.
    - 如果你的源文件包含后缀：_$GOOS.go 或 _$GOARCH.go，那么这个源文件只会在这个平台下编译. 可以连起来用 _$GOOS_$GOARCH.go（注意顺序).
    - 比如 mypkg_linux.go 、mypkg_windows_amd64.go 

## ref
1. [golang build 编译规则](https://www.cnblogs.com/hetonghai/p/6476510.html)