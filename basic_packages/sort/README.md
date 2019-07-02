# 简介
sort 是 golang 的标准库之一. 它提供了快速排序算法，可以对 Slice 进行排序及搜索. sort 源码总计 2500 多行. 除去测试代码仅 1100 行. 内部实现了快速排序、希尔排序、插入排序、堆排序等. 提供的 Sort 函数用于对切片进行快速排序.


## 类型
1. type interface
2. type IntSlice
3. type Float64Slice
4. type StringSlice

这些接口需要实现的方法包括
+ Len() int
+ Less(i,j int) bool
+ Search(x TYPE)
+ Sort()
+ Swap(i,j int)

## 公共函数
+ Ints(a []int)
+ IntsAreSorted(a []int) bool
+ SearchInts(a []int, x int) int
+ Float64s(a []float64)
+ Float64sAreSorted(a []float64) bool
+ SearchFloat64s(a []float64, x float64) int
+ Strings(a []string)
+ StringsAreSorted(a []string) bool
+ SearchStrings(a []string, x string) int
+ Sort(data Interface)
+ Stable(data Interface)
+ Reverse(data Interface) Interface
+ IsSorted(data Interface) bool
+ Search(n int, f func(int) bool) int

可以看到主要包括几个主要类型的排序、搜索、排序判断，以及对抽象接口的排序（不保证稳定性）、稳定排序、逆序、潘旭判断以及二分法搜索

## 内部主要函数
+ insertionSort()
+ siftDown()
+ heapSort()
+ medianOfThree()
+ swapRange()
+ doPivot()
+ quickSort()
+ maxDepth()
+ stable()
+ symMerge()
+ rotate()

内部函数主要涉及多个排序算法的实现

## Sort 快速排序
快速排序的内部实现中. 当分组长度大于 12 时，使用希尔排序. 否则使用插入排序.

    func quickSort(data Interface, a, b, maxDepth int) {
        for b-a > 12 { // Use ShellSort for slices <= 12 elements
            if maxDepth == 0 {
                heapSort(data, a, b)
                return
            }
            maxDepth--
            mlo, mhi := doPivot(data, a, b)
            // Avoiding recursion on the larger subproblem guarantees
            // a stack depth of at most lg(b-a).
            if mlo-a < b-mhi {
                quickSort(data, a, mlo, maxDepth)
                a = mhi // i.e., quickSort(data, mhi, b)
            } else {
                quickSort(data, mhi, b, maxDepth)
                b = mlo // i.e., quickSort(data, a, mlo)
            }
        }
        if b-a > 1 {
            // Do ShellSort pass with gap 6
            // It could be written in this simplified form cause b-a <= 12
            for i := a + 6; i < b; i++ {
                if data.Less(i, i-6) {
                    data.Swap(i, i-6)
                }
            }
            insertionSort(data, a, b)
        }
    }

golang 的排序对快速排序做了简单的优化. 可以思考下 12 为何可以作为分界点？