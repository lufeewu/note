/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	dp0, dp1 := cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		dp0, dp1 = dp1, min(dp0, dp1)+cost[i]
	}

	return min(dp0, dp1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

