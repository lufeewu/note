/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return [][]int{}
	}

	res = append(res, nums)
	next := nextPermutation(nums)
	for !numsEqual(nums, next) {
		res = append(res, next)
		next = nextPermutation(next)
	}
	return res
}

func numsEqual(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func nextPermutation(nums []int) []int {
	var res []int = make([]int, len(nums))
	copy(res, nums)
	n := len(res)
	i := n - 2
	for i >= 0 && res[i] >= res[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && res[i] >= res[j] {
			j--
		}
		res[i], res[j] = res[j], res[i]
	}

	reverse(res[i+1:])
	return res
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// @lc code=end

