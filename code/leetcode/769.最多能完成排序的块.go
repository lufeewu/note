/*
 * @lc app=leetcode.cn id=769 lang=golang
 *
 * [769] 最多能完成排序的块
 */

// @lc code=start
func maxChunksToSorted(arr []int) int {
	var res int
	m := 0
	for i := 0; i < len(arr); i++ {
		m = max(m, arr[i])
		if m == i {
			res++
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

