/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	courseLeft := numCourses

	for i := range prerequisites {
		from, to := prerequisites[i][1], prerequisites[i][0]
		graph[from] = append(graph[from], to)
		indeg[to]++
	}

	for courseLeft > 0 {
		idx := findZero(indeg)
		if idx == -1 {
			return false
		}

		courseLeft--
		for _, val := range graph[idx] {
			indeg[val]--
		}
	}

	return true
}

func findZero(nums []int) int {
	for i, val := range nums {
		if val == 0 {
			nums[i] = -1
			return i
		}
	}
	return -1
}

// @lc code=end

