/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	m, res, next := 0, 0, 0
	for i := 0; i < n-1; i++ {
		if m < nums[i]+i {
			m = nums[i] + i
		}
		if i == next {
			res++
			next = m
		}
	}
	return res
}

// @lc code=end

