/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	remain := 0
	curSum := 0
	res := 0
	for i := 0; i < len(gas); i++ {
		remain += gas[i] - cost[i]
		curSum += gas[i] - cost[i]
		if curSum < 0 {
			res = i + 1
			curSum = 0
		}
	}
	if remain < 0 {
		return -1
	}
	return res
}

// @lc code=end

