/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	var sum int = 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] = max(dp[j], dp[j-v]+v)
		}
	}

	return dp[target] == target
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

