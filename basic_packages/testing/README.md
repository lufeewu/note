# 简介
测试是 go 语言思想的一大特色，标准库 testing 为 go 语言提供自动化测试的支持.

## 单元测试
TestXxx(*testing.T)

## 基准测试
BenchmarkXxx(*testing.B)

## 示例验证测试
示例函数用于展示方法的使用，用于文档的效果. 另一方面，可以当做测试运行.
ExampleXxx_Xx_xxxx()

    func ExampleHello() {
        fmt.Println("Hello")
        // Output: Hello
    }

以 Example 开头的示例函数, 如果其中保护 "Output" 或 "Unordered output"  开头的注释，则在运行测试时，将函数的输出与注释进行比较.

## httptest

## 实现原理
+ godoc
+ go test
## ref 
1. [测试](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.0.html)
2. [go doc与godoc
](https://wiki.jikexueyuan.com/project/go-command-tutorial/0.5.html)