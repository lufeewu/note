/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
type SortBy [][]int

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}
	return a[i][0] > a[j][0]
}
func reconstructQueue(people [][]int) [][]int {
	sort.Sort(SortBy(people))
	for i := 1; i < len(people); i++ {
		cnt := 0
		for j := 0; j < i; j++ {
			if cnt == people[i][1] {
				tmp := people[i]
				for k := i - 1; k >= j; k-- {
					people[k+1] = people[k]
				}
				people[j] = tmp
				break
			}
			cnt++
		}
	}
	return people
}

// @lc code=end

