/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] 判断二分图
 */

// @lc code=start
var (
	UNCOLORED, RED, GREEN = 0, 1, 2
	color                 []int
	valid                 bool
)

func isBipartite(graph [][]int) bool {
	n := len(graph)
	valid = true
	color = make([]int, n)
	for i := 0; i < n && valid; i++ {
		if color[i] == UNCOLORED {
			dfs(i, RED, graph)
		}
	}
	return valid
}

func dfs(node, c int, graph [][]int) {
	color[node] = c
	cNei := RED
	if c == RED {
		cNei = GREEN
	}
	for _, neighbor := range graph[node] {
		if color[neighbor] == UNCOLORED {
			dfs(neighbor, cNei, graph)
			if !valid {
				return
			}
		} else if color[neighbor] != cNei {
			valid = false
			return
		}
	}
}

// @lc code=end

