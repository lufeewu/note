/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i, j := 1, 1

	for i < len(nums) && j < len(nums) {
		nums[i] = nums[j]
		i++
		j++
		for j < len(nums) && nums[j] == nums[i-1] &&
			nums[j] == nums[i-2] {
			j++
		}
	}

	return i
}

// @lc code=end

