/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	var r int = x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

// @lc code=end

