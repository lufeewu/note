/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	dp0, dp1 := 1, 1
	for i := 0; i < n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp0
}

// @lc code=end

