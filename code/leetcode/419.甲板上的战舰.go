/*
 * @lc app=leetcode.cn id=419 lang=golang
 *
 * [419] 甲板上的战舰
 */

// @lc code=start
func countBattleships(board [][]byte) int {
	lenthA := len(board)
	lenthB := len(board[0])
	count := 0
	for i := 0; i < lenthA; i++ {
		for j := 0; j < lenthB; j++ {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' || j > 0 && board[i][j-1] == 'X' {
					continue
				}
				count++
			}
		}
	}
	return count
}

// @lc code=end

