# 简介
标准库 html 提供了一些基本的用于转义和解转义 HTML 文本的函数.

## 源码
总计代码 11000 多行，除去测试代码总计 6700 多行.
+ entity.go 
    - 内部包括了所有 HTML entity 名称及值
+ escape.go 
    - 提供了转义与解转义 HTML entity 的函数
    - func UnescapeString(s string) string 
    - func EscapeString(s string) string
+ html/template 
    - 包实现了数据驱动的模板,可对抗代码注入.
    - 提供了同 text/template 相同的接口 
    - type Template struct 
        - 是 text/template 的特化版本，用于生成安全的 HTML 文本片段
        - func (t *Template) ParseFiles(filenames ...string) (*Template, error)
        - func (t *Template) Parse(src string) (*Template, error)
        - func Must(t *Template, err error) *Template
    - func (t *Template) Parse(src string) (*Template, error)
+ ...

## 应用
1. Gin 的 render 模块

## 其它同类库
1. text/template