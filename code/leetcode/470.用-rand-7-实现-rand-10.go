/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
func rand10() int {
	for {
		ans := (rand7()-1)*7 + rand7()
		if ans >= 1 && ans <= 40 {
			return (ans-1)%10 + 1
		}
	}
}

// @lc code=end

