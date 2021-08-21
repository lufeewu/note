/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	var left, right = 0, len(nums) - 1
	for {
		pos := quickSortHelper(nums, left, right)
		if pos == k-1 {
			return nums[pos]
		}
		if pos < k-1 {
			left = pos + 1
		}
		if pos > k-1 {
			right = pos - 1
		}
	}
	return -1
}

func quickSortHelper(nums []int, left int, right int) int {
	if left >= right {
		return left
	}
	pivot, l, r := nums[left], left, right
	for l < r {
		for l < r && nums[r] < pivot {
			r--
		}
		if l < r {
			nums[l] = nums[r]
			l++
		}
		for l < r && nums[l] > pivot {
			l++
		}
		if l < r {
			nums[r] = nums[l]
			r--
		}

	}
	nums[l] = pivot
	return l

}

// @lc code=end

