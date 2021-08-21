/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, k)
	var s []int
	for _, e := range nums {
		if _, ok := m[e]; !ok {
			s = append(s, e)
		}
		m[e]++
	}

	sort.Slice(s, func(i, j int) bool {
		return m[s[i]] > m[s[j]]
	})
	return s[:k]
}

// @lc code=end

