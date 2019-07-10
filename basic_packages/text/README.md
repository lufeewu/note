# 简介
text 标准库提供了文本相关的包，提供了字符编码、文本转换和特定区域设置的文本处理.它们与国际化(i18n)和本地化(l10)相关.

## 源码
总计 11600 多行代码, 除去测试代码仅 6254 多行.
+ text/scanner
    - 提供对 utf-8 文本的 token 扫描服务
    - type Position struct 表示一个位置
    - type Scanner struct 
        - 实现了读取 unicode 以及来自 io.Reader 的 token
        - func (s *Scanner) Scan() rune
        - ...
+ text/tabwriter
    - tabwriter 包实现了写入过滤包
    - 可以将输入的缩进修正为正确的对齐文本
+ text/template
    - template 包实现了数据驱动的用于生成文本输出的模板
    - HTML 格式输出接口同 html/template


## ref
1. [Elastic Tabstops](http://nickgravgaard.com/elastictabstops/index.html)