/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	var stack = make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if (s[i] == ')' && stack[len(stack)-1] != '(') ||
			(s[i] == ']' && stack[len(stack)-1] != '[') ||
			(s[i] == '}' && stack[len(stack)-1] != '{') {
			return false
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

// @lc code=end

