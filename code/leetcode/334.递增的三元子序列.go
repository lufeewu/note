/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}
	min1 := math.MaxInt32
	min2 := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if cur > min2 {
			return true
		} else if cur < min1 {
			min1 = cur
		} else if cur < min2 && cur > min1 {
			min2 = cur
		}
	}
	return false
}

// @lc code=end

