/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robHelper(nums[1:]), robHelper(nums[:len(nums)-1]))
}

func robHelper(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp1 := nums[0]
	dp2 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp1, dp2 = dp2, max(dp2, dp1+nums[i])
	}

	return dp2
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

