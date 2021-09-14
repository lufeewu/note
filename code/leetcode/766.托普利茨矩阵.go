/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		value := matrix[0][i]
		for j := 1; j < m && i+j < n; j++ {
			if matrix[j][i+j] != value {
				return false
			}
		}
	}

	for i := 1; i < m; i++ {
		value := matrix[i][0]
		for j := 1; j < n && i+j < m; j++ {
			if matrix[i+j][j] != value {
				return false
			}
		}
	}
	return true
}

// @lc code=end

