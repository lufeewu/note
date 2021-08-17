/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		minCoin := math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				minCoin = min(minCoin, dp[i-coin]+1)
			}
		}

		if minCoin == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = minCoin
		}
	}

	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

