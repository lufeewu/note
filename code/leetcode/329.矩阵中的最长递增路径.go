/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	M := make([][]int, m)
	for i := range M {
		M[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs(matrix, M, i, j, m, n))
		}
	}
	return res
}

func dfs(matrix [][]int, M [][]int, i, j, m, n int) int {
	if M[i][j] != 0 {
		return M[i][j]
	}
	// 默认长度为 1
	M[i][j] = 1
	// 上
	if j >= 1 && matrix[i][j-1] > matrix[i][j] {
		M[i][j] = max(M[i][j], 1+dfs(matrix, M, i, j-1, m, n))
	}
	// 下
	if j < n-1 && matrix[i][j+1] > matrix[i][j] {
		M[i][j] = max(M[i][j], 1+dfs(matrix, M, i, j+1, m, n))
	}
	// 左
	if i >= 1 && matrix[i-1][j] > matrix[i][j] {
		M[i][j] = max(M[i][j], 1+dfs(matrix, M, i-1, j, m, n))
	}
	// 右
	if i < m-1 && matrix[i+1][j] > matrix[i][j] {
		M[i][j] = max(M[i][j], 1+dfs(matrix, M, i+1, j, m, n))
	}

	return M[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

