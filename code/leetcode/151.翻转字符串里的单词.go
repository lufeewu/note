/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	b := []byte(s)
	var reverse func(start, end int)

	reverse = func(start, end int) {
		for start < end {
			b[start], b[end] = b[end], b[start]
			start += 1
			end -= 1
		}
	}
	reverse(0, len(b)-1)
	start := 0
	end := 0
	i := 0
	res := ""
	for i < len(b) {
		if b[i] != ' ' {
			start = i
			for i < len(b) && b[i] != ' ' {
				i += 1
			}
			end = i - 1
			reverse(start, end)
			res = res + " " + string(b[start:end+1])
		}
		i += 1
	}
	return res[1:]
}

// @lc code=end

