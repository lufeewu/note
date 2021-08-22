/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	cnt := 0
	var sum1Map = make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			sum1Map[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			if n1, ok := sum1Map[-(c + d)]; ok {
				cnt += n1
			}
		}
	}

	return cnt
}

// @lc code=end

