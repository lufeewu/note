/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	return subsetsWithDupDFS(nums)
}

func subsetsWithDupRange(nums []int) [][]int {
	var numCount = make(map[int]int)

	for _, n := range nums {
		numCount[n]++
	}

	var res [][]int
	res = append(res, []int{})

	for k, v := range numCount {
		for _, r := range res {
			for i := 1; i <= v; i++ {
				var tmp []int = make([]int, len(r))
				copy(tmp, r)
				for j := 0; j < i; j++ {
					tmp = append(tmp, k)
				}
				res = append(res, tmp)
			}
		}

	}

	return res
}

func subsetsWithDupDFS(nums []int) (ans [][]int) {
	sort.Ints(nums)
	t := []int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1]
	}
	dfs(false, 0)
	return
}

// @lc code=end

