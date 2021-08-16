/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var resLen int
	table := make(map[byte]int) //  记录字符上一次最新出现的位置, 从 1 开始
	subStart := 0               // 子串的开始下标
	for k, v := range s {
		if table[byte(v)] >= subStart { // v 已经出现过，更新子串开启的位置
			subStart = table[byte(v)] // 子串的下标从重复字符的下一个开始
		}
		table[byte(v)] = k + 1
		if resLen < (k - subStart + 1) {
			resLen = k - subStart + 1
		}
	}

	return resLen

}

// @lc code=end

