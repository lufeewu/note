/*
 * @lc app=leetcode.cn id=554 lang=golang
 *
 * [554] 砖墙
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	widthMap := make(map[int]int)
	maxWidth := 0
	for _, bricks := range wall {
		width := 0
		for i := 0; i < len(bricks)-1; i++ {
			width += bricks[i]
			if v, ok := widthMap[width]; ok {
				widthMap[width] = v + 1
			} else {
				widthMap[width] = 1
			}
			if maxWidth <= widthMap[width] {
				maxWidth = widthMap[width]
			}
		}
	}
	return len(wall) - maxWidth
}

// @lc code=end

