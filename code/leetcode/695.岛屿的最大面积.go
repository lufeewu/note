/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			area := dfs(grid, i, j)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) ||
		grid[i][j] == 0 {
		return 0
	}
	area := 1
	grid[i][j] = 0
	area += dfs(grid, i+1, j)
	area += dfs(grid, i-1, j)
	area += dfs(grid, i, j+1)
	area += dfs(grid, i, j-1)
	return area
}

// @lc code=end

