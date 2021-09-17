/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res, i, j := 0, 0, 0
	sum := nums[0]
	for i < len(nums) && j < len(nums) {
		if sum == target {
			if res == 0 || res > j-i+1 {
				res = j - i + 1
			}
			sum = sum - nums[i]
			i++
			j++
			if j < len(nums) {
				sum = sum + nums[j]
			}
		} else if sum > target {
			if res == 0 || res > j-i+1 {
				res = j - i + 1
			}
			sum = sum - nums[i]
			i++
		} else {
			j++
			if j < len(nums) {
				sum = sum + nums[j]
			}
		}
	}
	return res
}

// @lc code=end

