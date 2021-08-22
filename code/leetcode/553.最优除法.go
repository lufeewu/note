/*
 * @lc app=leetcode.cn id=553 lang=golang
 *
 * [553] 最优除法
 */

// @lc code=start
func optimalDivision(nums []int) string {
	var sli []string

	for _, n := range nums {
		sli = append(sli, strconv.Itoa(n))
	}

	if len(sli) < 3 {
		return strings.Join(sli, "/")
	}

	return sli[0] + "/(" + strings.Join(sli[1:], "/") + ")"
}

// @lc code=end

