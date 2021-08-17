/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	if s < S || (s+S)%2 == 1 {
		return 0
	}

	target := (s + S) / 2       // 数组中和为 target 存在的数量，s - 2target = S
	dp := make([]int, target+1) // 和为 i 的数量
	dp[0] = 1

	for _, v := range nums {
		for i := target; i >= v; i-- {
			dp[i] = dp[i] + dp[i-v]
		}
	}

	return dp[target]
}

// @lc code=end

