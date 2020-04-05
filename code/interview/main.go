package main

import "fmt"

func main() {

	fmt.Println(SplitNumbers([]int{6, 4, -3, 0, 5, -2, -1, 0, 1, -9}))
	fmt.Println(SplitNumbers([]int{-9, 1, 0, -1, -2, 5, 0, -3, 4, 6}))
	fmt.Println(SplitNumbers([]int{0}))
	fmt.Println(SplitNumbers([]int{-1, -1}))
	fmt.Println(SplitNumbers([]int{1, 3}))
	fmt.Println(SplitNumbers([]int{}))
	fmt.Println(SplitNumbers([]int{2, 3, -2, 0, 0}))
	fmt.Println(SplitNumbers([]int{0, 0, 0}))
	fmt.Println(SplitNumbers([]int{0, 0, 0, -1, 3, 4, 5, 6}))

	fmt.Println(SerializeReversed(map[string]string{"1": "bar",
		"2": "foo.bar",
		"3": "foo.bar.cloud",
		"4": "baz.cloudmall.com",
		"5": "baz.cloudmall.ai",
	}))
	fmt.Println(SerializeReversed(map[string]string{}))
	fmt.Println(SerializeReversed(nil))
	fmt.Println(SerializeReversed(map[string]string{"1": "bar",
		"2": "foo.bar",
		"3": "foo.bar.cloud",
		"4": "baz.cloudmall.com",
		"5": "baz.cloudmall.ai",
	}))
}
