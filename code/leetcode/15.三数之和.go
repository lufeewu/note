/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	ans := [][]int{}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1
		target := -nums[i]

		for l < r {
			if nums[l]+nums[r] == target {
				ans = append(ans, []int{nums[l], nums[r], nums[i]})
				left, right := nums[l], nums[r]
				for r >= 0 && nums[r] == right {
					r--
				}
				for l < n && nums[l] == left {
					l++
				}
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}

// threeSum 方法, 排序后查找 O(n2logn), 使用 map
func threeSumQuery(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	var res [][]int
	table := make(map[int]int)
	for k, v := range nums {
		table[v] = k
	}

	// 查询
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if v, ok := table[-nums[i]-nums[j]]; ok {
				if v > i && v > j {
					res = append(res, []int{nums[i], nums[j], nums[v]})
				}
			}
		}
	}
	return res
}

// @lc code=end

