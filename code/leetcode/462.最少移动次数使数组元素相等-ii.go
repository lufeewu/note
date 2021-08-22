/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最少移动次数使数组元素相等 II
 */

// @lc code=start
func minMoves2(nums []int) int {
	lenth := len(nums)
	sort.Ints(nums)
	num := nums[lenth/2]
	res := 0
	for i := 0; i < lenth; i++ {
		res += getAbs(num, nums[i])
	}
	return res
}

// 中位数性质推导: https://zhuanlan.zhihu.com/p/73139689

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// @lc code=end

