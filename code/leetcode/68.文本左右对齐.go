/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) (ans []string) {
	right, n := 0, len(words)
	for {
		left := right // 当前行的第一个单词在 words 的位置
		sumLen := 0   // 统计这一行单词长度之和
		// 循环确定当前行可以放多少单词，注意单词之间应至少有一个空格
		for right < n && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		// 当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格
		if right == n {
			s := strings.Join(words[left:], " ")
			ans = append(ans, s+blank(maxWidth-len(s)))
			return
		}

		numWords := right - left
		numSpaces := maxWidth - sumLen

		// 当前行只有一个单词：该单词左对齐，在行末填充剩余空格
		if numWords == 1 {
			ans = append(ans, words[left]+blank(numSpaces))
			continue
		}

		// 当前行不只一个单词
		avgSpaces := numSpaces / (numWords - 1)
		extraSpaces := numSpaces % (numWords - 1)
		s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1)) // 拼接额外加一个空格的单词
		s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))  // 拼接其余单词
		ans = append(ans, s1+blank(avgSpaces)+s2)
	}
}

// blank 返回长度为 n 的由空格组成的字符串
func blank(n int) string {
	return strings.Repeat(" ", n)
}

// @lc code=end

