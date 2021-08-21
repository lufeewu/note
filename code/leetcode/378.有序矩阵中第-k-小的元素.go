/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])

	lo := matrix[0][0]
	hi := matrix[m-1][n-1] + 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		j := n - 1
		i := 0
		for ; i < m; i++ {
			for j >= 0 && mid < matrix[i][j] {
				j--
			}
			count += j + 1
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// @lc code=end

