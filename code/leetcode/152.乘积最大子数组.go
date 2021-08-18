/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, maxV, minV := nums[0], nums[0], nums[0]

	for _, v := range nums[1:] {
		v1, v2 := v*maxV, v*minV
		maxV = max(v, v1, v2)
		minV = min(v, v1, v2)
		res = max(res, maxV)
	}
	return res
}

func max(a int, args ...int) int {
	for _, v := range args {
		if a < v {
			a = v
		}
	}
	return a
}

func min(a int, args ...int) int {
	for _, v := range args {
		if a > v {
			a = v
		}
	}
	return a
}

// @lc code=end

