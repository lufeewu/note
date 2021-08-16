/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	var res int32
	var i uint
	// 计算每个 bit 位出现的次数
	for ; i < 32; i++ {
		var cnt uint
		for k := 0; k < len(nums); k++ {
			cnt += (uint(nums[k]) >> i) & 1 // 第 i bit 位出现的次数
		}

		// 比特位上出现次数不是 3 的倍数的位数
		if cnt%3 != 0 {
			res = res | (1 << i)
		}
	}
	return int(res)
}

// @lc code=end

