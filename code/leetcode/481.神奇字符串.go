/*
 * @lc app=leetcode.cn id=481 lang=golang
 *
 * [481] 神奇字符串
 */

// @lc code=start
func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	s := make([]byte, 0, n+2)
	s = append(s, '1', '2', '2')
	next := 2
	res := 1
	for i := 2; i < n; i++ {
		if s[i] == '1' {
			res++
			if s[next] == '2' {
				s = append(s, '1')
			} else {
				s = append(s, '2')
			}
		} else {
			if s[next] == '2' {
				s = append(s, '1', '1')
			} else {
				s = append(s, '2', '2')
			}
			next++
		}
		next++
	}

	return res
}

// @lc code=end

