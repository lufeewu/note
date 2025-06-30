package main

import (
	"fmt"
	"sync"

	"github.com/lufeewu/note/code/sort"
)

func max(first int, args ...int) int {
	var res = first
	for _, v := range args {
		if v > res {
			res = v
		}
	}
	return res
}

func f(n int) int {
	f1, f2 := 1, 1
	for i := 0; i < n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f1
}

var wg sync.WaitGroup
var slic1 []int
var lock sync.Mutex

func f2(i int) {
	defer wg.Done()
	// lock.Lock()
	slic1 = append(slic1, i)
	// lock.Unlock()
	fmt.Println(i, len(slic1), cap(slic1))
}

// 并发安全
func testA() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f2(i)
	}
	wg.Wait()
	fmt.Println(slic1)
	return
}

var wg1 sync.WaitGroup

func a() {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}

}
func b() {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}

}
func testB() {
	// fmt.Println(runtime.NumCPU())
	// runtime.GOMAXPROCS(12)
	wg1.Add(2)
	go a()
	// wg1.Add(1)

	go b()
	wg1.Wait()
	// time.Sleep(10 * time.Second)
}

func canPartition(nums []int) bool {
    sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum % 2 != 0 {
		return false
	}
	dp := make([]int, sum / 2 + 1)
	for _, v := range nums {
		// 按背包从大到小放入, 不会出现背包放放两次情况
		for i := len(dp) - 1; i >= v ; i-- {
			if dp[i] < dp[i-v] + v {
				dp[i] = dp[i-v] + v
			} 
		}
	}
	return dp[len(dp) - 1] == sum/2
}

func main() {
	fmt.Println(canPartition([]int{1, 2, 5}))
	return

	fmt.Println(f(0), f(1), f(2), f(3), f(4), f(5))

	fmt.Println(max(3, 4, 5, 6))
	return
	list := []int{10, 9, 8, 7, -1, 5, 4, 3, 2, 1}
	sort.SelectSort(list)
	sort.InsertSort(list)
	sort.BubbleSort(list)
	sort.QuickSort(list, 0, len(list))
	sort.ShellSort(list)
	sort.HeapSort(list)
}
