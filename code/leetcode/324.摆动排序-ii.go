/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 */

// @lc code=start
func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	sort.Ints(nums)
	mid := len(nums)/2 + len(nums)%2
	var res []int

	for i := 0; i < mid; i++ {
		res = append(res, nums[mid-i-1])
		if mid+i < len(nums) {
			res = append(res, nums[len(nums)-i-1])
		}
	}
	for i := 0; i < len(res); i++ {
		nums[i] = res[i]
	}
}

// @lc code=end

