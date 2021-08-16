/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n > 0 {
		return fastPow(x, n)
	}
	return 1.0 / fastPow(x, -n)
}

func fastPow(x float64, n int) float64 {
	var res float64 = 1.0
	for n != 0 {
		if n%2 == 1 {
			res = res * x
		}
		x = x * x
		n = n >> 1
	}
	return res
}

// @lc code=end

