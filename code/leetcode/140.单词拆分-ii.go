/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	ret := make([]string, 0)
	m := make(map[string]bool, 0)
	for _, w := range wordDict {
		m[w] = true
	}

	path := make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			ret = append(ret, strings.Join(path, " "))
			return
		}
		for i := start + 1; i <= len(s); i++ {
			if m[s[start:i]] {
				path = append(path, s[start:i])
				dfs(i)
				path = path[0 : len(path)-1]
			}
		}
	}
	dfs(0)
	return ret
}

// @lc code=end

