/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return 0
	}
	a, sum := 0, 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			a = a + 1 // a表示以nums[i-1]结尾的等差数组个数。如 [1,2,3,4]，以4结尾的等差数组有2个[1,2,3,4]和[2,3,4]，加上5的话，有3个：[1,2,3,4,5]、[2,3,4,5]和[3,4,5]
		} else {
			a = 0
		}
		sum += a
	}
	return sum
}

// @lc code=end

