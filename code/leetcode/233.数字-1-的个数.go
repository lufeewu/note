/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 */

// @lc code=start
func countDigitOne(n int) int {
	var res, digit int

	// 计算各个数位上为 1 时，出现的次数，然后计算出现在 1 的总数量和
	for i := 1; i <= n; i *= 10 {
		res += n / (i * 10) * i // 某一位数上是 1, 左边的组合数(不含最大)*右边组合的总数
		digit = (n / i) % 10

		// 左边取最大值情况
		if digit > 1 { // 该位大于 1 , 则取 1 后, 右边的组合可以任意取
			res += i
		} else if digit == 1 { // 该位为 1 , 则右边数量有限, 即 n % i + 1
			res += n%i + 1
		} // 该位为 0

	}
	return res
}

// @lc code=end

