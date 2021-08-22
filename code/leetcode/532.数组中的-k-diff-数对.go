/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的 k-diff 数对
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	if n < 2 {
		return 0
	}
	var res int
	i, j := 0, 1
	for i < n && j < n {
		for j < n-1 && nums[j+1] == nums[j] {
			j++
		}
		for i < n-1 && i < j && nums[i+1] == nums[i] {
			i++
		}
		if nums[j]-nums[i] == k {
			res++
			j++
		} else if nums[j]-nums[i] > k {
			i++
		} else {
			j++
		}
		if j == i {
			j = i + 1
		}

	}
	return res
}

// @lc code=end

