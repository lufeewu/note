/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp := make([]int, n)  // 记录最长子序列长度
	cnt := make([]int, n) // 记录下标 i 后的子序列的最长递增子序列个数
	l, ans := 1, 0
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		cnt[i] = 1
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = 1 + dp[j]
					cnt[i] = cnt[j]
				} else if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j]
				}
			}
		}
		if l == dp[i] {
			ans += cnt[i]
		} else if l < dp[i] {
			l = dp[i]
			ans = cnt[i]
		}
	}

	return ans
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

