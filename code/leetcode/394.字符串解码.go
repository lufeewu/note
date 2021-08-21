/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	var stack []string
	var tmp string
	var prv byte

	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '[':
			stack = append(stack, tmp)
			stack = append(stack, "[")
			tmp = ""
		case s[i] == ']':
			if len(tmp) > 0 {
				stack = append(stack, tmp)
			}
			repeat, rpStr := 0, ""
			s := stack[len(stack)-1]
			for s != "[" {
				rpStr = s + rpStr
				stack = stack[:len(stack)-1]
				s = stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			repeat, _ = strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack = append(stack, strings.Repeat(rpStr, repeat))

			tmp = ""
		case isDigital(s[i]):
			if isDigital(prv) {
				tmp += string(s[i])
			} else if isAlpha(prv) {
				stack = append(stack, tmp)
				tmp = string(s[i])
			} else {
				tmp = string(s[i])
			}
		case isAlpha(s[i]):
			tmp += string(s[i])
			if i == len(s)-1 {
				stack = append(stack, string(tmp))
			}
		}
		prv = s[i]
		// fmt.Println(stack)
	}
	res := ""
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}

	return res
}

func isDigital(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// @lc code=end

