/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, left int, right int) {
	for i := left; i < (right+1+left)/2; i++ {
		nums[i], nums[right+left-i] = nums[right+left-i], nums[i]
	}
}

// @lc code=end

