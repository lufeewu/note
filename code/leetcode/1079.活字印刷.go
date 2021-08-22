/*
 * @lc app=leetcode.cn id=1079 lang=golang
 *
 * [1079] 活字印刷
 */

// @lc code=start
func numTilePossibilities(tiles string) int {
	if len(tiles) == 0 {
		return 0
	}
	var array = make([]int, 26)
	for _, c := range tiles {
		array[c-'A']++
	}
	return DFS(array)
}

func DFS(array []int) int {
	var sum = 0
	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			continue
		}
		sum++
		array[i]--
		sum += DFS(array)
		array[i]++
	}
	return sum
}

// @lc code=end

