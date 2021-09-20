/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	var res []byte

	tmp := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) && a[len(a)-1-i] == '1' {
			tmp++
		}
		if i < len(b) && b[len(b)-i-1] == '1' {
			tmp++
		}
		res = append(res, byte('0'+tmp%2))
		tmp = tmp / 2
	}
	if tmp > 0 {
		res = append(res, byte('0'+tmp))
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return string(res)
}

// @lc code=end

