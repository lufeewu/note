/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 */

// @lc code=start
func lastRemaining(n int) int {
	return HandlN(n, 0)
}

func HandlN(n int, cnt int) int {
	if n == 1 {
		return 1
	}
	if cnt%2 == 0 || n%2 == 1 {
		return 2 * HandlN(n>>1, cnt+1)
	}
	return 2*HandlN(n>>1, cnt+1) - 1
}

// @lc code=end

