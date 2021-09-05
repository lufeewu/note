# 简介
debug 是 golang 标准库中，用于提供程序运行时 debug 的工具库.

## 源码
总计 15000 多行代码, 除去测试代码 11000 多行.
+ dwarf
    - 提供从可执行文件中加载 DWARF 格式的 debug 信息
+ elf
    - 可以访问 ELF 格式的文件
+ gosym
    - 实现了用于访问 Go 符号和行数表
    - 由 gc 编译器生成的嵌入到 Go 二进制中
+ macho
    - 实现了用于访问 Mach-O 对象文件
+ pe
    - 实现了用于访问 PE(MS Portable Executable) 格式文件
+ plan9obj
    - 实现了用于访问 Plan 9 a.out 对象文件

## 应用
1. 调试工具 github.com/go-delve/delve

## ref
1. [DWARF Debugging Information Format](http://dwarfstd.org/doc/dwarf-2.0.0.pdf)