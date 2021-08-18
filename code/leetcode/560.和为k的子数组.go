/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	var res, sum = 0, 0
	n := len(nums)
	table := make(map[int]int) // 前面 n 个数和对应的数量
	table[0] = 1
	for i := 0; i < n; i++ {
		sum += nums[i]
		res += table[sum-k]
		table[sum]++
	}
	return res
}

// @lc code=end

