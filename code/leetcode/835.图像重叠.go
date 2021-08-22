/*
 * @lc app=leetcode.cn id=835 lang=golang
 *
 * [835] 图像重叠
 */

// @lc code=start
func largestOverlap(A [][]int, B [][]int) int {
	var res, n = 0, len(A)

	var listA, listB [][]int
	var diffCnt = make(map[string]int, 1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				listA = append(listA, []int{i, j})
			}
			if B[i][j] == 1 {
				listB = append(listB, []int{i, j})
			}
		}
	}

	for _, a := range listA {
		for _, b := range listB {
			diffCnt[fmt.Sprintf("%d-%d", a[0]-b[0], a[1]-b[1])]++
		}
	}

	for _, d := range diffCnt {
		if res < d {
			res = d
		}
	}
	return res
}

// @lc code=end

