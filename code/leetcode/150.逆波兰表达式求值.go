/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := list.New()
	for _, token := range tokens {
		switch token {
		case "+":
			op2 := stack.Back()
			stack.Remove(op2)
			op1 := stack.Back()
			stack.Remove(op1)
			stack.PushBack(op1.Value.(int) + op2.Value.(int))
		case "-":
			op2 := stack.Back()
			stack.Remove(op2)
			op1 := stack.Back()
			stack.Remove(op1)
			stack.PushBack(op1.Value.(int) - op2.Value.(int))
		case "*":
			op2 := stack.Back()
			stack.Remove(op2)
			op1 := stack.Back()
			stack.Remove(op1)
			stack.PushBack(op1.Value.(int) * op2.Value.(int))
		case "/":
			op2 := stack.Back()
			stack.Remove(op2)
			op1 := stack.Back()
			stack.Remove(op1)
			stack.PushBack(op1.Value.(int) / op2.Value.(int))
		default:
			el, _ := strconv.Atoi(token)
			stack.PushBack(el)
		}
	}
	res := stack.Front()
	return res.Value.(int)
}

// @lc code=end

