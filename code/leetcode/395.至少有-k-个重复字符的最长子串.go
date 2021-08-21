/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	if k == 1 {
		return len(s)
	}
	if k > len(s) {
		return 0
	}

	mem := make(map[rune]int)
	for _, v := range s {
		mem[v]++
	}

	max := 0
	currentOK := true
	leftBound := 0
	for i, v := range s {
		if mem[v] < k {
			if i-leftBound >= k {
				if tmp := longestSubstring(s[leftBound:i], k); tmp > max {
					max = tmp
				}
			}
			leftBound = i + 1
			currentOK = false
		}
	}

	// process last substring
	if !currentOK {
		if len(s)-leftBound >= k {
			if tmp := longestSubstring(s[leftBound:], k); tmp > max {
				max = tmp
			}
		}
	}

	if currentOK {
		return len(s)
	} else {
		return max
	}
}

// @lc code=end

