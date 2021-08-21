/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start

func hIndex(A []int) int {
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		if A[i] >= len(A)-i {
			return len(A) - i
		}
	}
	return 0
}

// @lc code=end

