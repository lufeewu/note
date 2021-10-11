/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}
	nums = nums[:i+1]
	return len(nums)
}

// @lc code=end

