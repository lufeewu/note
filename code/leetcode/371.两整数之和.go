/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	// 异或+与运算:时间复杂度O(logSum) | 空间复杂度O(1)
	a, b = a^b, (a&b)<<1
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

// @lc code=end

