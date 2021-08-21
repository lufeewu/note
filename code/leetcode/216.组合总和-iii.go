/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var dfsHelper = &DFSHelper{
		K: k,
		N: n,
	}
	dfsHelper.DFS(1, n)
	return dfsHelper.Res
}

type DFSHelper struct {
	Item []int
	Res  [][]int
	K    int
	N    int
}

func (d *DFSHelper) DFS(cur, rest int) {
	// 找到一个答案
	if len(d.Item) == d.K && rest == 0 {
		d.Res = append(d.Res, append([]int(nil), d.Item...))
		return
	}
	// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
	if len(d.Item)+10-cur < d.K || rest < 0 {
		return
	}
	// 跳过当前数字
	d.DFS(cur+1, rest)
	// 选当前数字
	d.Item = append(d.Item, cur)
	d.DFS(cur+1, rest-cur)
	d.Item = d.Item[:len(d.Item)-1]
}

// @lc code=end

