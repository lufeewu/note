/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	var parents = make([]int, len(edges))
	for i, _ := range parents {
		parents[i] = i + 1
	}

	for _, e := range edges {
		if union(parents, e[0], e[1]) == false {
			return e
		}
	}

	return []int{}

}

func find(parents []int, x int) int {
	if parents[x-1] == x {
		return x
	}
	return find(parents, parents[x-1])
}

func union(parents []int, from, to int) bool {
	x, y := find(parents, from), find(parents, to)
	if x == y {
		return false // uinon 失败, from to 已经是连通的
	}
	parents[x-1] = y
	return true
}

// @lc code=end

