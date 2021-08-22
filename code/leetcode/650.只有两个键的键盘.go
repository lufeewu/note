/*
 * @lc app=leetcode.cn id=650 lang=golang
 *
 * [650] 只有两个键的键盘
 */

// @lc code=start
func minSteps(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			res += i
			n /= i
		}
	}
	return res
}

// @lc code=end

