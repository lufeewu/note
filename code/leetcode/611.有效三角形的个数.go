/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var res int
	// 遍历第三条边，找前两条边之和大于第三条边的组合
	for k := len(nums) - 1; k >= 2; k-- {
		l, r := 0, k-1
		for l < r {
			if nums[l]+nums[r] > nums[k] {
				res += r - l
				r--
			} else {
				l++
			}
		}
	}
	return res
}

// @lc code=end

