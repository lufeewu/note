/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) int {
	var table = make(map[int]int)
	table[0] = -1
	res, sum := 0, 0
	for i := 0; i < len(nums); i++ {

		if nums[i] == 0 {
			sum += -1
		} else {
			sum += 1
		}

		if v, ok := table[sum]; ok {
			res = max(i-v, res)
		} else {
			table[sum] = i
		}
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

