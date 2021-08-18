/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationSumHelper(candidates, target)
}

func combinationSumHelper(candidates []int, target int) (res [][]int) {
	for k, v := range candidates {
		if v > target {
			break
		}
		if v == target {
			res = append(res, []int{v})
			break
		}

		next := candidates[k:len(candidates)]
		tmp := combinationSum(next, target-v)
		for _, t := range tmp {
			tmpRes := []int{v}
			tmpRes = append(tmpRes, t...)
			res = append(res, tmpRes)
		}
	}
	return res
}

// @lc code=end

