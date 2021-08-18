/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	seq := "1"
	for j := 1; j < n; j++ {
		i := 0
		nextSeq := ""
		for i < len(seq) {
			count := 1
			for i < len(seq)-1 && seq[i] == seq[i+1] {
				count++
				i++
			}
			nextSeq += strconv.Itoa(count) + string(seq[i])
			i++
		}
		seq = nextSeq
	}
	return seq
}

// @lc code=end

