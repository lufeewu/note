/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}

		prev := res[len(res)-1]
		if prev[1] < interval[0] {
			res = append(res, interval)
			continue
		}

		prev[1] = max(prev[1], interval[1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

