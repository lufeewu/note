/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 */

// @lc code=start
func isMonotonic(nums []int) bool {
	inc, dec := true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			inc = false
		}
		if nums[i] < nums[i+1] {
			dec = false
		}
	}
	return inc || dec

}

// @lc code=end

