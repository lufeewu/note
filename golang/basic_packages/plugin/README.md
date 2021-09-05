# 简介
plugin 是 golang 的标准库，目前 plugin 仅在 linux 上有效. 它可以让程序在运行时动态加载外部的功能.

## 源码
总计 261 行代码，除去测试代码仅包括 243 行代码
+ type Plugin struct
    - func Open(path string) (*Plugin, error)
    - func (p *Plugin) Lookup(symName string) (Symbol, error)
+ type Symbol interface

## 应用
1. docker containerd 模块使用了 plugin，用于加载 OS 和 Arch 的所有 plugins


## ref
1. [Golang笔记-Plugin初探](https://www.jianshu.com/p/4ab799081a99)