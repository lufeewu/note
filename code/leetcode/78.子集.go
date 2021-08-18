/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	return dfs(0, nums, []int{})
}

func dfs(i int, nums []int, list []int) [][]int {
	var res [][]int
	if i == len(nums) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		return res
	}
	r := dfs(i+1, nums, list)
	if len(r) > 0 {
		res = append(res, r...)
	}
	list = append(list, nums[i])
	r = dfs(i+1, nums, list)
	if len(r) > 0 {
		res = append(res, r...)
	}

	return res
}

// @lc code=end

