package main

import "fmt"

// 2. slice 与 append
func sliceAppend() {
	// slice 在 append 时候，只有当触发扩容的时候会复制一个新的数组，否则将会用原数组
	// 当两个变量指向同一个 slice 地址时候，将会有数据被覆盖的风险
	var src, dest []int
	fmt.Println(cap(src))
	src = make([]int, 0, 12)
	src = append(src, []int{1, 2, 3, 4}...)
	dest = src
	fmt.Printf("%p %p %v %v \n", src, dest, src, dest)
	dest = append(dest, 5)
	fmt.Printf("%p %p %v %v \n", src, dest, src, dest)
	src = append(src, 6)
	fmt.Printf("%p %p %v %v \n", src, dest, src, dest) // src 与 dest 的地址一样, dest 的数据 5 被覆盖为 6
	src = append(src, 7)
	fmt.Printf("%p %p %v %v \n", src, dest, src, dest)
	dest = append(dest, 8)
	fmt.Printf("%p %p %v %v \n", src, dest, src, dest) // src 与 dest 的地址一样, src 的数据 7 被覆盖为 8

}
