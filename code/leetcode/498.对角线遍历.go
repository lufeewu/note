/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var res, temp []int
	row, col, flag, i, j := len(matrix), len(matrix[0]), 0, 0, 0
	if row == 1 {
		return matrix[0]
	}
	for j <= col-1 {
		m, n := i, j
		temp = append(temp, matrix[i][j])
		for m != 0 && n != col-1 {
			m--
			n++
			temp = append(temp, matrix[m][n])
		}
		if flag%2 != 0 {
			reverse(temp)
		}
		res = append(res, temp...)
		flag++
		temp = temp[:0]
		if i != row-1 {
			i++
		} else {
			j++
		}
	}
	return res
}

func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// @lc code=end

