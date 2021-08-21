/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {
	ans, cur := 0, 0
	stack := []int{} // 里面的数全部相加

	preOp := -1 // -1:+ -2:- -3:* -4:/

	for i, ch := range s {
		// num
		n := getNum(byte(ch))
		if n >= 0 {
			cur = cur*10 + n
		}
		if (n < 0 && n > -5) || i == len(s)-1 { // end of a num
			if preOp == -1 { // +
				stack = append(stack, cur)
			} else if preOp == -2 { // -
				stack = append(stack, -cur)
			} else if preOp == -3 { // *
				tmp := stack[len(stack)-1] * cur
				stack[len(stack)-1] = tmp
			} else if preOp == -4 { // /
				tmp := stack[len(stack)-1] / cur
				stack[len(stack)-1] = tmp
			}
			preOp = n
			cur = 0
		}
	}

	for _, n := range stack {
		ans += n
	}
	return ans
}

func getNum(char byte) int {
	if char-'0' >= 0 && char-'0' <= 9 {
		return int(char - '0')
	} else if char == '+' {
		return -1
	} else if char == '-' {
		return -2
	} else if char == '*' {
		return -3
	} else if char == '/' {
		return -4
	}
	return -5
}

// @lc code=end

