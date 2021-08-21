/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	var res int
	for i := len(nums) - 1; i >= 0; i-- {
		tmp := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] && dp[j] > tmp {
				tmp = dp[j]
			}
		}
		dp[i] = 1 + tmp
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// @lc code=end

