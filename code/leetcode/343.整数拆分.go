/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	var dp = make([]int, n+2)
	for i := 2; i <= n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*(i-j), j*dp[i-j])
		}
	}

	return dp[n]
}

func max(a, b, c int) int {
	if b > a && b > c {
		return b
	} else if c > b && c > a {
		return c
	}
	return a
}

// @lc code=end

