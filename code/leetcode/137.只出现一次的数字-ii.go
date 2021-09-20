/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	return singleNumberStateHelper(nums)
}

func singleNumberStateHelper(nums []int) int {
	one, two := 0, 0
	for _, num := range nums {
		one = (one ^ num) & ^two
		two = (two ^ num) & ^one
	}
	return one
}

func singleNumberBitsHelper(nums []int) int {
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

func singleNumberQuickSortHelper(nums []int, left, right int) int {
	if right-left == 0 {
		return nums[left]
	}
	k := rand.Intn(right + 1 - left)
	pivot := nums[left+k]
	nums[left], nums[k] = nums[k], nums[left]
	l, r := left, right
	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] < pivot {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	if l > left && (l-left)%3 > 0 {
		return singleNumberQuickSortHelper(nums, left, l-1)
	}
	return singleNumberQuickSortHelper(nums, l, right)
}

// @lc code=end

