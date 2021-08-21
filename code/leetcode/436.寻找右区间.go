/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] 寻找右区间
 */

// @lc code=start
func findRightInterval(intervals [][]int) []int {
	var rightMap = make(map[int][]int)
	var res, left []int
	for i, interval := range intervals {
		rightMap[interval[0]] = append(rightMap[interval[0]], i)
		left = append(left, interval[0])
	}
	sort.Ints(left)

	for _, interval := range intervals {
		rightIndex := findRightIndex(left, interval[1])
		if rightIndex == -1 || rightIndex >= len(left) {
			res = append(res, -1)
			continue
		}
		if k, ok := rightMap[left[rightIndex]]; ok {
			res = append(res, k[0])
		} else {
			res = append(res, -1)
		}

	}

	return res
}

func findRightIndex(nums []int, value int) int {
	if value > nums[len(nums)-1] {
		return -1
	}
	l, r := 0, len(nums)-1
	mid := (l + r) / 2
	for l < r {
		if nums[mid] < value {
			l = mid + 1
		} else {
			r = mid
		}
		mid = (l + r) / 2
	}

	return l
}

// @lc code=end

