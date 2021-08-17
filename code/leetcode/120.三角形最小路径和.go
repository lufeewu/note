/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	var dp []int
	for _, v := range triangle[len(triangle)-1] {
		dp = append(dp, v)
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for k, v := range triangle[i] {
			dp[k] = min(v+dp[k], v+dp[k+1])
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

