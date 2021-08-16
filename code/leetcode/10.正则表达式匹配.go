/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	// 表示输入串长度为 i，模式串长度为 j 时, 是否匹配
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	if m > 0 && n > 0 {
		dp[1][1] = s[0] == p[0] || p[0] == '.'
	}
	for j := 2; j <= m; j++ {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	for i := 1; i <= n; i += 1 {
		for j := 2; j <= m; j += 1 {
			if p[j-1] != '*' { // 新增的字符必须和当前 s 的末尾匹配
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			} else { // 最后一个是 *，则末尾可以是 0 个或多个匹配
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
				// 若为多个匹配，则当前 s 少了后，首先是可以匹配的，而新增的字符与前一个相同或者是任意的符合要求
			}
		}
	}
	return dp[n][m]
}

// @lc code=end

