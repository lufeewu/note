/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) (ans int) {
	visited := make([]bool, len(isConnected))
	for i, v := range visited {
		if !v {
			ans++
			dfs(isConnected, visited, i)
		}
	}
	return ans
}

func dfs(isConnected [][]int, visited []bool, from int) {
	visited[from] = true
	for to, conn := range isConnected[from] {
		if conn == 1 && !visited[to] {
			dfs(isConnected, visited, to)
		}
	}
}

// @lc code=end

