/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {

	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[l] > nums[mid] {
			l++
			r = mid
		} else {
			return nums[l]
		}
	}
	return nums[l]

}

// @lc code=end

