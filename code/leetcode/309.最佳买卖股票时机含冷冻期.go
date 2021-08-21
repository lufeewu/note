/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	sell := make([]int, len(prices))
	buy := make([]int, len(prices))
	sell[0] = 0
	buy[0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]) // cooldown or sell
		if i > 1 {
			buy[i] = max(buy[i-1], sell[i-2]-prices[i]) // cooldown or buy
		} else {
			buy[i] = max(buy[i-1], -prices[i]) // cooldown or buy
		}
	}
	return sell[len(prices)-1]
}

func max(first int, args ...int) int {
	var res = first
	for _, v := range args {
		if v > res {
			res = v
		}
	}
	return res
}

// @lc code=end

