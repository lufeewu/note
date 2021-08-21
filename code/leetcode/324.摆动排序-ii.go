/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 */

// @lc code=start
func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	sort.Ints(nums)

	for i := 1; i < len(nums); i = i + 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
}

// @lc code=end

