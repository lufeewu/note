/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	helper(n, n, "", &res)
	return res
}

func helper(l, r int, tmp string, res *[]string) {
	if r < l {
		return
	}
	if l == 0 && r == 0 {
		*res = append(*res, tmp)
	}
	if l > 0 {
		helper(l-1, r, tmp+"(", res)
	}
	if r > 0 {
		helper(l, r-1, tmp+")", res)
	}
}

// @lc code=end

