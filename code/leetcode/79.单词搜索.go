/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == word[0] {
				used := make([]int, h*w)
				if existHelper(board, used, word, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func existHelper(board [][]byte, used []int, word string, i int, j int) bool {
	if word == "" {
		return true
	}
	h, w := len(board), len(board[0])
	if i < 0 || i >= h || j < 0 || j >= w {
		return false
	}

	if used[i*w+j] == 1 {
		return false
	}

	if word[0] != board[i][j] {
		return false
	}

	used[i*w+j] = 1
	if existHelper(board, used, word[1:], i+1, j) ||
		existHelper(board, used, word[1:], i, j+1) ||
		existHelper(board, used, word[1:], i-1, j) ||
		existHelper(board, used, word[1:], i, j-1) {
		return true
	}
	used[i*w+j] = 0
	return false

}

// @lc code=end

