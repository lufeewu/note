/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	var n1, n2, n = len(num1), len(num2), len(num1)
	if n2 > n {
		n = n2
	}
	var res = make([]byte, 0, n)
	var tmp = 0
	for i := 0; i < n; i++ {
		if i < n1 {
			tmp += int(num1[n1-i-1] - '0')
		}
		if i < n2 {
			tmp += int(num2[n2-i-1] - '0')
		}
		res = append(res, byte('0'+tmp%10))
		tmp = tmp / 10
	}
	for tmp > 0 {
		res = append(res, byte('0'+tmp%10))
		tmp = tmp / 10
	}

	// reverse
	n = len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return string(res)
}

// @lc code=end

