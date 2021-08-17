/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	visited := make([]bool, len(graph))
	res := [][]int{}

	visited[0] = true
	route := []int{0}
	res = dfs(visited, graph, route, 0)
	return res
}

func dfs(visited []bool, graph [][]int, route []int, k int) [][]int {
	var res [][]int
	if k == len(graph)-1 {
		res = append(res, append([]int{}, route...))
		return res
	}
	for i := 0; i < len(graph[k]); i++ {
		if !visited[graph[k][i]] {
			visited[graph[k][i]] = true

			route = append(route, graph[k][i])
			r := dfs(visited, graph, route, graph[k][i])
			if len(r) > 0 {
				res = append(res, r...)
			}
			route = route[:len(route)-1]

			visited[graph[k][i]] = false
		}
	}

	return res
}

// @lc code=end

