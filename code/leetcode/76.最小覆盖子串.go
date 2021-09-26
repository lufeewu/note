/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	var res string
	m := make(map[byte]int) // 剩余需要的次数，负数代表窗口中的冗余

	for _, v := range t {
		m[byte(v)]++
	}
	left, cnt, minLen := 0, 0, len(s)+1

	for i := 0; i < len(s); i++ {
		m[s[i]] -= 1
		if m[s[i]] >= 0 {
			cnt++
		}

		for cnt == len(t) {
			if minLen > i-left+1 {
				minLen = i - left + 1
				res = s[left : left+minLen]
			}
			m[s[left]] += 1
			if m[s[left]] > 0 { // 大于 0 的，多余部分是 t 中出现的
				// 仅在 s 中的字符 value 最多是 0
				cnt--
			}
			left++
		}
	}

	return res
}

// @lc code=end

