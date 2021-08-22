/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	var m = make([]int, 26)

	for _, t := range tasks {
		m[t-'A']++
	}

	sort.Ints(m)

	i := 25
	for ; i >= 0 && m[25] == m[i]; i-- {

	}
	return max((m[25]-1)*(n+1)+25-i, len(tasks))
}

func max(nums ...int) int {
	ret := 0
	for _, num := range nums {
		if num > ret {
			ret = num
		}
	}
	return ret
}

// @lc code=end

