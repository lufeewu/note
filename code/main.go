package main

import (
	"github.com/lufeewu/note/code/sort"
)

func main() {
	list := []int{10, 9, 8, 7, -1, 5, 4, 3, 2, 1}
	sort.SelectSort(list)
	sort.InsertSort(list)
	sort.BubbleSort(list)
	sort.QuickSort(list)
	sort.ShellSort(list)
	sort.HeapSort(list)
}