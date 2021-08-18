/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var res int
	for i < j {
		if height[i] < height[j] {
			res = max(res, (j-i)*height[i])
			i++
		} else {
			res = max(res, (j-i)*height[j])
			j--
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

