/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start

func generateMatrix(n int) [][]int {
	var DIRS = [][]int{[]int{0, 1}, []int{1, 0},
		[]int{0, -1}, []int{-1, 0}} // 右下左上
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	r, c, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[r][c] = i
		dir := DIRS[dirIdx]
		r, c = r+dir[0], c+dir[1]
		if r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			r, c = r-dir[0], c-dir[1]
			dirIdx = (dirIdx + 1) % 4 // 顺时针旋转至下一个方向
			dir = DIRS[dirIdx]
			r, c = r+dir[0], c+dir[1]
		}
	}
	return matrix
}

// @lc code=end

