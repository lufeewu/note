/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m - 1; i >= 0; i-- {
		nums1[n+i] = nums1[i]
	}

	i, j := n, 0
	for i+j < m+2*n {
		if i >= m+n {
			nums1[i-n+j] = nums2[j]
			j++
		} else if j >= n {
			nums1[i-n+j] = nums1[i]
			i++
		} else if nums1[i] <= nums2[j] {
			nums1[i-n+j] = nums1[i]
			i++
		} else {
			nums1[i-n+j] = nums2[j]
			j++
		}
	}
}

// @lc code=end

