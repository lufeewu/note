/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	return updateMatrixDP(mat)
}
func updateMatrixDFS(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return [][]int{}
	}
	var res = make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[0]))
		for j := 0; j < len(mat[0]); j++ {
			res[i][j] = -1
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			dfs(mat, i, j, res)
		}
	}
	return res
}

func dfs(mat [][]int, i, j int, res [][]int) {
	if i >= len(mat) || j >= len(mat[0]) {
		return
	}
	if mat[i][j] == 0 {
		res[i][j] = 0
		return
	}

	// 向右递归
	if j+1 < len(mat[0]) && res[i][j+1] < 0 {
		dfs(mat, i, j+1, res)
	}
	// 向下递归
	if i+1 < len(mat) && res[i+1][j] < 0 {
		dfs(mat, i+1, j, res)
	}

	// 向下寻找
	if j+1 < len(mat[0]) && res[i][j+1] >= 0 {
		res[i][j] = res[i][j+1] + 1
	}

	// 向右寻找
	if i+1 < len(mat) && res[i+1][j] >= 0 &&
		(res[i][j] < 0 || res[i+1][j] < res[i][j]) {
		res[i][j] = res[i+1][j] + 1
	}

	// 向上寻找
	if i-1 >= 0 && res[i-1][j] >= 0 &&
		(res[i][j] < 0 || res[i-1][j] < res[i][j]) {
		res[i][j] = res[i-1][j] + 1
	}

	// 向左寻找
	if j-1 >= 0 && res[i][j-1] >= 0 &&
		(res[i][j] < 0 || res[i][j-1] < res[i][j]) {
		res[i][j] = res[i][j-1] + 1
	}
}

func updateMatrixDP(mat [][]int) [][]int {
	dp := make([][]int, len(mat))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(mat[0]))
		for j := 0; j < len(mat[0]); j++ {
			dp[i][j] = 100010
		}
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if i > 0 {
					dp[i][j] = Min(dp[i][j], dp[i-1][j]+1)
				}
				if j > 0 {
					dp[i][j] = Min(dp[i][j], dp[i][j-1]+1)
				}
			}
		}
	}

	for i := len(dp) - 1; i >= 0; i-- {
		for j := len(dp[0]) - 1; j >= 0; j-- {
			if i < len(dp)-1 {
				dp[i][j] = Min(dp[i][j], dp[i+1][j]+1)
			}
			if j < len(dp[0])-1 {
				dp[i][j] = Min(dp[i][j], dp[i][j+1]+1)
			}
		}
	}

	return dp
}

func Min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

// @lc code=end

