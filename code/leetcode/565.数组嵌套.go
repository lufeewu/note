/*
 * @lc app=leetcode.cn id=565 lang=golang
 *
 * [565] 数组嵌套
 */

// @lc code=start
func arrayNesting(nums []int) int {
	// 设置一个指针
	max := 0
	for i := range nums { // 遍历这个array
		p := 1 // 设置一个每次的转换次数的初始值
		for nums[i] != i {
			// 看题意  A[0], A[5], A[6], A[2]} 当 a[2] == 0的时候就退出了，
			// 0刚好是  a[0] == 5的index，所以只要 nums[i] == i 就该跳出了。
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			p++
		}
		if max < p {
			max = p
		}
	}
	return max
}

// @lc code=end

