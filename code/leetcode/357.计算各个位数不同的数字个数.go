/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {

	dp := make([]int, 1)
	dp[0] = 1 //起点
	//dp数组中下标代表n位时的个数

	for i := 1; i <= n; i++ {
		sum := 9 //sum值默认为9 因为一位时是9种

		for j := i - 1; j > 0; j-- {
			sum *= 10 - j
		}

		temp := sum
		temp += dp[i-1] //n=2时，dp[2]=两位数的个数 加上一位数的个数 也就是等于dp[1]+两位数的个数
		dp = append(dp, temp)
	}
	return dp[n]
}

// @lc code=end

