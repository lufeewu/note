/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */

// @lc code=start
func findLongestChain(pairs [][]int) int {
	arr := Pairs{}
	arr = pairs
	sort.Sort(arr)

	cnt, max := 0, 1<<31-1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i][len(arr[i])-1] < max {
			max = arr[i][0]
			cnt++
		}
	}

	return cnt
}

type Pairs [][]int

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	if p[i][0] < p[j][0] {
		return true
	}
	if p[i][0] == p[j][0] && p[i][len(p[i])-1] < p[j][len(p[j])-1] {
		return true
	}
	return false
}
func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// @lc code=end

