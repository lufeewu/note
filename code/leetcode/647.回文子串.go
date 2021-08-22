/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		helper(s, i, i, &count)   // odd
		helper(s, i, i+1, &count) // even
	}
	return count
}

func helper(s string, left, right int, c *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		*c++
		left--
		right++
	}
}

// @lc code=end

