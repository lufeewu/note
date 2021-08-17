/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}

	for _, e := range edges {
		if !union(parent, e[0], e[1]) {
			return e
		}
	}
	return nil
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent []int, from, to int) bool {
	x, y := find(parent, from), find(parent, to)
	if x == y {
		return false
	}
	parent[x] = y
	return true
}

// @lc code=end

