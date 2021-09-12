/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	if len(words) < 1 {
		return []int{}
	}

	wlen := len(words)
	k := len(words[0])
	if len(s) < k*wlen {
		return []int{}
	}

	var res []int
	var mp = make(map[string]int)
	for _, w := range words {
		mp[w]++
	}

	for i := 0; i < len(s)-k*wlen+1; i++ {
		var count int
		var mp2 = make(map[string]int)
		for multi := 0; multi < wlen; multi++ {
			start := i + multi*k
			word := s[start : start+k]
			if num, found := mp[word]; found && num > mp2[word] {
				mp2[word]++
				count++
			} else {
				break
			}
		}
		if count == wlen {
			res = append(res, i)
		}
	}

	return res
}

// @lc code=end

