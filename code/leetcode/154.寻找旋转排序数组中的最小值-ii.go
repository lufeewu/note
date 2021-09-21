/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {

	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] < nums[r] {
			return nums[l]
		}
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[l] > nums[mid] {
			l++
			r = mid
		} else {
			r--
		}
	}
	return nums[l]
}

// @lc code=end

