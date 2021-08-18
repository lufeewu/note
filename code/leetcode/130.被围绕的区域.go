/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte) {
	n := len(board)
	if n == 0 {
		return
	}

	m := len(board[0])
	if m == 0 {
		return
	}

	// first and last row
	for j := 0; j < m; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j, n, m)
		}
		if board[n-1][j] == 'O' {
			dfs(board, n-1, j, n, m)
		}
	}

	for i := 1; i < n-1; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0, n, m)
		}
		if board[i][m-1] == 'O' {
			dfs(board, i, m-1, n, m)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'F' {
				board[i][j] = 'O'
			}
		}
	}

}

func dfs(board [][]byte, i int, j int, n int, m int) {
	board[i][j] = 'F'
	if i+1 < n && board[i+1][j] == 'O' {
		dfs(board, i+1, j, n, m)
	}
	if i-1 >= 0 && board[i-1][j] == 'O' {
		dfs(board, i-1, j, n, m)
	}
	if j+1 < m && board[i][j+1] == 'O' {
		dfs(board, i, j+1, n, m)
	}
	if j-1 >= 0 && board[i][j-1] == 'O' {
		dfs(board, i, j-1, n, m)
	}
}

// @lc code=end

