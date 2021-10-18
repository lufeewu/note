# 简介
golang 的 gc 主要是基于标记-清扫(mark and sweep)算法. 经典的 gc 算法有三种: 引用计数(reference counting)、标记-清扫(mark & sweep)、复制收集(Copy and Collection).

## 引用计数(reference counting)
通过对被引用对象进行计数引用，当计数器为 0 时，释放内存。

## 标记-清扫(mark & sweep)
标记清扫可能会造成 STW, golang 优化了标记清扫，通过三色标记法减少 STW 带来的问题。

## 复制收集(Copy and Collection)
将内存容量划分为大小相等的两块，每次只使用其中一块。当这一块的内存用完了，就将还存活着的对象复制到另外一块上面，然后把已经使用过的内存空间一次清理掉。

## 参考
