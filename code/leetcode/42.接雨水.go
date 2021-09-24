/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	var res, max int
	var maxStack = make([]int, len(height))
	maxStack[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > maxStack[i-1] {
			maxStack[i] = height[i]
		} else {
			maxStack[i] = maxStack[i-1]
		}
	}

	max = height[len(height)-1]
	for i := len(height) - 2; i > 0; i-- {
		if height[i] >= max {
			max = height[i]
		} else if maxStack[i-1] > max {
			res += max - height[i]
		} else if maxStack[i-1] > height[i] {
			res += maxStack[i-1] - height[i]
		}
	}
	return res
}

// @lc code=end

