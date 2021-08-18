/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	cache := make(map[int]int)
	cache[0] = 1
	return dfs(s, cache)
}

func dfs(s string, cache map[int]int) int {
	if v, ok := cache[len(s)]; ok {
		return v
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		cache[1] = 1
		return 1
	}

	if s[0] == '1' || (s[0] == '2' && s[1] >= '0' && s[1] <= '6') {
		n1 := dfs(s[1:], cache)
		n2 := dfs(s[2:], cache)
		cache[len(s)-1] = n1
		cache[len(s)-2] = n2
		return n1 + n2
	}
	n := dfs(s[1:], cache)
	cache[len(s)-1] = n
	return n
}

// @lc code=end

