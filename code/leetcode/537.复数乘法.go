/*
 * @lc app=leetcode.cn id=537 lang=golang
 *
 * [537] 复数乘法
 */

// @lc code=start
func complexNumberMultiply(a string, b string) string {
	indexAPlus := strings.Index(a, "+")
	indexAI := strings.Index(a, "i")
	indexBPlus := strings.Index(b, "+")
	indexBI := strings.Index(b, "i")
	x, _ := strconv.Atoi(a[:indexAPlus])
	y, _ := strconv.Atoi(a[indexAPlus+1 : indexAI])
	c, _ := strconv.Atoi(b[:indexBPlus])
	d, _ := strconv.Atoi(b[indexBPlus+1 : indexBI])
	return fmt.Sprint(x*c-y*d, "+", y*c+x*d, "i")
}

// @lc code=end

