/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n < 1 {
		return 0
	}
	f0, f1 := 0, 1
	for n != 0 {
		f0, f1 = f1, f0+f1
		n--
	}

	return f0
}

// @lc code=end

