package main

import (
	"fmt"
	"sync"
)

// 1. 值拷贝、for 值、go 协程、函数闭包
func forgo() {
	var wg sync.WaitGroup
	wg.Add(12)
	for i := 0; i < 6; i++ { // i 最终的值是 6
		go func() {
			fmt.Print(i, ",") // 是 for 中的一个变量，它的地址不会变
			wg.Done()
		}()
	}
	for i := 0; i < 6; i++ {
		go func(i int) {
			fmt.Print(i, ",") // i 是函数参数
			wg.Done()
		}(i) // 参数传递，值拷贝
	}
	wg.Wait()
}
