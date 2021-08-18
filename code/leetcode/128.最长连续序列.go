/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	ret := 0
	m := map[int]int{}
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			l := m[v-1]
			r := m[v+1]
			sum := l + r + 1
			m[v] = sum

			ret = max(ret, sum)
			// 改变当前连续数的边界的计数值即可，只有它们将会影响后面的值的计算
			m[v-l] = sum
			m[v+r] = sum
		} else {
			continue
		}
	}
	return ret
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

