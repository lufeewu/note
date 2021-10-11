/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start

func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--

	ans := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k/factorial[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n-i]
	}
	return ans
}

func getPermutationNextPermutation(n int, k int) string {
	if n < 2 {
		return "1"
	}
	var nums = make([]byte, n)
	for i := 0; i < n; i++ {
		nums[i] = byte('0' + i + 1)
	}

	for i := 0; i < k-1; i++ {
		nextPermutation(nums)
	}
	return string(nums)
}

func nextPermutation(nums []byte) {
	if len(nums) < 2 {
		return
	}
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 {
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}
	if i < 0 {
		reverse(nums)
		return
	}

	for i < j {
		if nums[j] > nums[i] {
			break
		}
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
	return
}

func reverse(nums []byte) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

// @lc code=end

