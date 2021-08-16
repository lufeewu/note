/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	var count = 0
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}

// @lc code=end

