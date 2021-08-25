/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	if n < 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ { // top is j
			dp[i] += dp[j-1] * dp[i-j] // 左子树数量 * 右子树数量
		}
	}
	return dp[n]
}

// @lc code=end

