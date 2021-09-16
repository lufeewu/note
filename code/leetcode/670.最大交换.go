/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	if num < 10 {
		return num
	}
	var arr []int = make([]int, 0, 10)
	var stack []int

	var n = num
	for n != 0 {
		arr = append(arr, n%10)
		n = n / 10
	}

	n = -1
	for i := 0; i < len(arr)-1; i++ {
		if n < 0 || arr[i] > arr[n] {
			n = i
		}
		stack = append(stack, n)
	}
	for i := len(arr) - 1; i >= 1; i-- {
		if arr[i] < arr[stack[i-1]] {
			arr[i], arr[stack[i-1]] = arr[stack[i-1]], arr[i]
			break
		}
	}
	n = 0
	for i := 0; i < len(arr); i++ {
		n = n*10 + arr[len(arr)-i-1]
	}
	return n
}

// @lc code=end

