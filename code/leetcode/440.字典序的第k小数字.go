/*
 * @lc app=leetcode.cn id=440 lang=golang
 *
 * [440] 字典序的第K小数字
 */

// @lc code=start
func findKthNumber(n int, k int) int {
	var p, prefix = 1, 1
	for p < k {
		count := getCount(prefix, n)
		if p+count > k {
			/// 说明第k个数，在这个前缀范围里面
			prefix *= 10
			p++
		} else if p+count <= k {
			/// 说明第k个数，不在这个前缀范围里面，前缀需要扩大+1
			prefix++
			p += count
		}
	}
	return prefix
}

func getCount(prefix, n int) int {
	var cur = prefix
	var next, count int = cur + 1, 0

	for cur <= n {
		count += min(n+1, next) - cur
		cur *= 10
		next *= 10
	}
	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

