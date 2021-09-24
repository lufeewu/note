/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[l] < nums[mid] && target < nums[mid] && target >= nums[l] {
			r = mid - 1
		} else if nums[mid] < nums[r] && target > nums[mid] && target <= nums[r] {
			l = mid + 1
		} else if nums[l] > nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return -1
}

// @lc code=end

