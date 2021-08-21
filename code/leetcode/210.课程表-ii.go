/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses) // 矩阵记录有向图
	indeg := make([]int, numCourses)   // 记录入度

	// Initialize the graph and indeg.
	for _, prereq := range prerequisites {
		from, to := prereq[1], prereq[0]
		graph[from] = append(graph[from], to)
		indeg[to]++
	}

	var res []int

	// Top sort.
	for len(res) < numCourses {
		i := getZeroIdx(indeg)
		if i == -1 {
			return []int{}
		}

		res = append(res, i) // 拓扑顺序

		for _, val := range graph[i] {
			indeg[val]-- // 所有 i 开始的路径，终点入度减 1
		}
	}
	return res
}

func getZeroIdx(nums []int) int {
	for i, val := range nums {
		if val == 0 {
			nums[i] = -1
			return i
		}
	}
	return -1 // 没有找到入度有 0 的边，存在环
}

// @lc code=end

