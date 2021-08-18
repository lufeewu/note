/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])

	// first row and first column
	r0 := 1
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			r0 = 0
			break
		}
	}

	// zero record
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// set zero depend 0 column
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// set zero depend 0 row
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < m; j++ {
				matrix[j][i] = 0
			}
		}
	}

	// set first row
	if r0 == 0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

}

// @lc code=end

