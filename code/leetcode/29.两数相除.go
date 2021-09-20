/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := 0
	for dividend >= divisor {
		tmp, m := divisor, 1
		for dividend >= tmp<<1 {
			tmp = tmp << 1
			m = m << 1
		}
		dividend -= tmp
		res += m
	}
	return sign * res
}

// @lc code=end

