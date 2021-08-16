/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res, tmp int
	res = nums[0]
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		if tmp > res {
			res = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return res
}

// @lc code=end

