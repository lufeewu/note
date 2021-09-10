/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) (ans [][]int) {
	return combinationSum2DP(candidates, target)
}
func combinationSum2DFS(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func combinationSum2DP(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}

	sort.Ints(candidates)
	var dp = make([][][]int, target+1)
	for _, v := range candidates {
		if v > target {
			break
		}
		for j := target; j >= v+1; j-- {
			if len(dp[j-v]) > 0 {
				for _, arr := range dp[j-v] {
					var tmp = make([]int, len(arr))
					copy(tmp, arr)
					tmp = append(tmp, v)
					dp[j] = append(dp[j], tmp)
					dp[j] = removeDuplicate(dp[j])
				}
			}
		}

		dp[v] = append(dp[v], []int{v})
		dp[v] = removeDuplicate(dp[v])
	}

	res = dp[target]
	return res
}

// 顺序去重
func removeDuplicate(arr [][]int) [][]int {
	if len(arr) == 0 {
		return arr
	}
	var res [][]int
	for i := 0; i < len(arr); i++ {
		var exist = false
		for j := 0; j < len(res); j++ {
			if sameOrderArr(arr[i], arr[j]) {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, arr[i])
		}
	}
	return res
}

// 顺序数组比较
func sameOrderArr(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// @lc code=end

