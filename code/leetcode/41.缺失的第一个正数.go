/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {

	j := len(nums) - 1
	i := 0

	for i < j {
		if nums[i] == i+1 {
			i++
			continue
		}
		if nums[i] < j+1 && nums[i] > i && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
		if nums[i] == nums[j] {
			j--
		}
	}
	i = 0
	for i < len(nums) {
		if nums[i] != i+1 {
			return i + 1
		}
		i++
	}
	return i + 1
}

// @lc code=end

