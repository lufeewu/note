/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	dp := make([][]bool, len(s)) // 记录从 i 到 j 的子串是否为回文
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true // 单个字符的字符串是回文
	}
	start := 0
	end := 0
	for j := 1; j < len(s); j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > end-start+1 {
					start = i
					end = j
				}
			}
		}
	}

	return s[start : end+1]
}

// @lc code=end

