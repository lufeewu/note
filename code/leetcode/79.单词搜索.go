/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {

	words, m, n := []byte(word), len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == words[0] { // 寻找到匹配的第一个字符
				if search(i, j, board, words) {
					return true
				}
			}
		}
	}

	return false
}

func search(i, j int, board [][]byte, words []byte) bool {

	m, n := len(board), len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != words[0] {
		// 不符合的条件，索引越界和字符不等
		return false
	}

	if len(words) == 1 {
		// 匹配成功， 返回true
		return true
	}

	tmp := board[i][j]
	board[i][j] = '1' // 由于words只能是字母，所以'1'不会被匹配

	// 递归回溯
	if search(i+1, j, board, words[1:]) ||
		search(i, j+1, board, words[1:]) ||
		search(i-1, j, board, words[1:]) ||
		search(i, j-1, board, words[1:]) {
		return true
	} else {
		//注意由于board是slice引用类型，所以函数的修改会真正的修改原slice的值，所以需要重新改正回来
		board[i][j] = tmp
		return false
	}
}

// @lc code=end

