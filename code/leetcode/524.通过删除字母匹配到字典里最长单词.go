/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
func findLongestWord(s string, dictionary []string) string {
	var temp string
	sort.Strings(dictionary)
	for _, i := range dictionary {
		if isMatch(s, i) == true && len(i) > len(temp) {
			temp = i
		}
	}
	return temp
}

func isMatch(s string, sub string) bool {
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sub[j] {
			j++
		}
		if j == len(sub) {
			return true
		}
	}
	return false
}

// @lc code=end

