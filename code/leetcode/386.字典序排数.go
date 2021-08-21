/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start
func lexicalOrder(n int) []int {
	ret := make([]int, 0)
	num := 1
	for {
		if num <= n {
			ret = append(ret, num)
			num *= 10
		} else {
			num /= 10
			for num%10 == 9 {
				num /= 10
			}
			if num == 0 {
				break
			}
			num++
		}
	}
	return ret
}

// @lc code=end

