/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */

// @lc code=start
func smallestRangeI(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	if max-min <= 2*k {
		return 0
	}

	return max - min - 2*k

}

// @lc code=end

