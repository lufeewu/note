/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	// QuickSort(nums, 0, len(nums)-1)
	// sort.Ints(nums)
	HeapSort(nums)
	return nums
}

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot, l, r := nums[left], left, right
	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		if l < r {
			nums[l] = nums[r]
		}
		for l < r && nums[l] <= pivot {
			l++
		}
		if l < r {
			nums[r] = nums[l]
		}

	}
	nums[l] = pivot
	QuickSort(nums, left, l-1)
	QuickSort(nums, l+1, right)
}

// HeapSort 堆排序
func HeapSort(nums []int) {
	n := len(nums)
	// 建堆, 最大堆
	for i := n/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, n)
	}
	// 取出数据放到末尾, 并调整堆
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}

}

func adjustHeap(nums []int, l, r int) {
	k := l
	for 2*k+1 < r {
		j := 2*k + 1
		if j+1 < r && nums[j] < nums[j+1] {
			j++
		}
		if nums[k] < nums[j] {
			nums[k], nums[j] = nums[j], nums[k]
			k = j
		} else {
			break
		}
	}
}

// @lc code=end

