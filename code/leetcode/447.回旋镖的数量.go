/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int { // 两重循环，hash 表
	ans := 0
	hash := map[int]int{}
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i != j {
				d := dist(points[i], points[j])
				if hash[d] > 0 { // 这个距离出现过，累计回旋镖个数
					ans += hash[d] * 2
				}
				hash[d]++
			}
		}
		hash = map[int]int{} // 算新的起点时，清空 hash 表
	}
	return ans
}
func dist(i, j []int) int { // 两点间的距离
	return (i[0]-j[0])*(i[0]-j[0]) + (i[1]-j[1])*(i[1]-j[1])
}

// @lc code=end

