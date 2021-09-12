/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}
	dfs(0, line, column, block, spaces, board)
}

func dfs(pos int, line, column [9][9]bool, blocks [3][3][9]bool,
	spaces [][2]int, board [][]byte) bool {
	if pos > len(spaces) {
		return false
	}
	if pos == len(spaces) {
		return true
	}

	i, j := spaces[pos][0], spaces[pos][1]
	for digit := byte(0); digit < 9; digit++ {
		if !line[i][digit] && !column[j][digit] &&
			!blocks[i/3][j/3][digit] {
			line[i][digit] = true
			column[j][digit] = true
			blocks[i/3][j/3][digit] = true
			board[i][j] = digit + '1'
			if dfs(pos+1, line, column, blocks, spaces, board) {
				return true
			}
			line[i][digit] = false
			column[j][digit] = false
			blocks[i/3][j/3][digit] = false
		}

	}
	return false

}

// @lc code=end

