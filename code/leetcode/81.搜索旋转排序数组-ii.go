/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target || nums[l] == target || nums[r] == target {
			return true
		} else if nums[l] < target && target < nums[mid] {
			r = mid - 1
			l++
		} else if nums[mid] < target && target < nums[r] {
			l = mid + 1
			r--
		} else if nums[l] > nums[mid] {
			r = mid - 1
			l++
		} else if nums[r] < nums[mid] {
			l = mid + 1
			r--
		} else {
			l++
			r--
		}
	}
	return false
}

// @lc code=end

