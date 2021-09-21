/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	var res int64
	var flag = true
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	first := parseNum(s[0])

	if first == -3 {
		return 0
	} else if first == -2 {
		flag = true
	} else if first == -1 {
		flag = false
	} else {
		res = int64(first)
	}

	for _, s := range s[1:] {
		if s == ' ' {
			break
		}
		n := parseNum(byte(s))
		if n < 0 {
			break
		} else {
			res = 10*res + int64(n)
			if res > 1<<32 {
				break
			}
		}
	}

	if flag == false {
		res = -res
	}
	if res < -(1 << 31) {
		return -(1 << 31)
	}
	if res >= 1<<31 {
		return 1<<31 - 1
	}
	return int(res)
}

func parseNum(s byte) int {
	if s <= '9' && s >= '0' {
		return int(s - '0')
	}
	if s == '-' {
		return -1
	}
	if s == '+' {
		return -2
	}
	return -3
}

// @lc code=end

