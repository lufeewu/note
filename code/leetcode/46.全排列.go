/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var res [][]int

	res = append(res, nums)
	next := nextPermutation(nums)
	for !equalArr(nums, next) {
		res = append(res, next)
		next = nextPermutation(next)
	}
	return res

}

func nextPermutation(nums []int) []int {
	var res []int = make([]int, len(nums))
	copy(res, nums)
	if len(nums) < 2 {
		return res
	}
	i := len(res) - 2
	for i >= 0 && res[i] <= res[i+1] {
		i--
	}
	if i >= 0 {
		j := len(res) - 1
		for j > i && res[i] <= res[j] {
			j--
		}

		res[i], res[j] = res[j], res[i]
	}

	reverse(res[i+1:])
	return res
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

func equalArr(arr1, arr2 []int) bool {
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

// @lc code=end

