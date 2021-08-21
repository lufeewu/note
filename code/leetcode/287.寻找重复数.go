/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	l, r, mid := 0, len(nums)-1, 0

	for l < r {
		mid = (l + r) / 2
		c := count(nums, mid)
		if c <= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func count(nums []int, mid int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= mid {
			c++
		}
	}
	return c
}

// @lc code=end

