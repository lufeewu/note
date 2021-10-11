/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) (ans [][]int) {
	var tmp []int
	combineHelper(n, k, 1, &tmp, &ans)
	return ans
}

func combineHelper(n int, k int, cur int, tmp *[]int, ans *[][]int) {
	// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
	if len(*tmp)+(n-cur+1) < k {
		return
	}
	// 记录合法的答案
	if len(*tmp) == k {
		comb := make([]int, k)
		copy(comb, *tmp)
		*ans = append(*ans, comb)
		return
	}

	// 考虑选择当前位置
	*tmp = append(*tmp, cur)

	combineHelper(n, k, cur+1, tmp, ans)

	*tmp = (*tmp)[:len(*tmp)-1]
	// 考虑不选择当前位置
	combineHelper(n, k, cur+1, tmp, ans)
}

// @lc code=end

