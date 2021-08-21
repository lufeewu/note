/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	var data []int
	n := len(nums)
	for i := 0; i < n; i++ {
		index := nums[i]
		if index < 0 {
			index = -index
		}
		if nums[index-1] < 0 {
			data = append(data, index)
		} else {
			nums[index-1] = -nums[index-1]
		}
	}
	return data
}

// @lc code=end

