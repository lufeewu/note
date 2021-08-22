/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	// 记录 A 中下标 i 开头、B 中下标 j 开头的重复子数组的长度
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

// @lc code=end

