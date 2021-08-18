/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if (mid-1 < 0 || nums[mid] > nums[mid-1]) &&
			(mid+1 == len(nums) || nums[mid] > nums[mid+1]) {
			return mid
		} else if (mid-1 < 0 || nums[mid] >= nums[mid-1]) &&
			(mid+1 == len(nums) || nums[mid] <= nums[mid+1]) { //上升
			left = mid + 1
		} else { //下降
			right = mid - 1
		}
	}

	return 0
}

// @lc code=end

