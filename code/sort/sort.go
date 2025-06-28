package sort

/*
 1. 简单排序
	- 选择排序
	- 插入排序
	- 冒泡排序
*/
// 可参考文章: https://www.cnblogs.com/gaopeng527/p/6699648.html

// SelectSort 选择排序
func SelectSort(list []int) {
	for i, _ := range list {
		tmp := i
		for j := i; j < len(list); j++ {
			if list[tmp] > list[j] {
				tmp = j
			}
		}
		list[tmp], list[i] = list[i], list[tmp]
	}
	return
}

// InsertSort 插入排序
func InsertSort(list []int) {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			} else {
				break
			}
		}
	}
	return
}

// BubbleSort 冒泡排序
func BubbleSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 1; j < len(list)-i; j++ {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	return
}

/*
  2. 归并排序
*/
func merge(list []int, a, b int, tmp []int) {
	var (
		i, j int
		mid  int
		pos  int
	)

	copy(tmp, list)
	i = a
	mid = (a + b) / 2
	j = mid + 1
	pos = a
	for i <= mid && j <= b {
		if tmp[i] <= tmp[j] {
			list[pos] = tmp[i]
			pos++
			i++
		} else {
			list[pos] = tmp[j]
			pos++
			j++
		}
	}
	for i <= mid {
		list[pos] = tmp[i]
		pos++
		i++
	}
	for j <= b {
		list[pos] = tmp[j]
		pos++
		j++
	}
}

// MergeSort 归并排序左闭右闭[a,b]
func MergeSort(list []int, a, b int) {
	if a+1 >= b {
		if list[a] > list[b] {
			list[a], list[b] = list[b], list[a]
		}
		return
	}
	tmp := make([]int, len(list))
	// 排序包括下标(a+b)/2的值
	MergeSort(list, a, (a+b)/2)

	MergeSort(list, (a+b)/2+1, b)

	merge(list, a, b, tmp)
}

// QuickSort 快速排序
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

// QuickSort2 快速排序实现 2
func QuickSort2(list []int) {
	var (
		midValue   int
		head, tail int
	)
	if len(list) <= 1 {
		return
	}
	head = 1
	midValue = list[0]
	tail = len(list) - 1
	for head < tail {
		for list[head] < midValue && head < tail {
			head++
		}
		if head >= tail {
			break
		}
		for list[tail] > midValue && tail > head {
			tail--
		}
		if tail <= head {
			break
		}
		list[head], list[tail] = list[tail], list[head]
		head++
		tail--
	}

	if list[head] > midValue {
		list[0], list[head-1] = list[head-1], list[0]

		QuickSort2(list[0 : head-1])
		QuickSort2(list[head:])
	} else {
		list[0], list[head] = list[head], list[0]

		QuickSort2(list[0:head])
		QuickSort2(list[head+1:])
	}

}

/*HeapSort 堆排序
  4. 堆排序
	 nlog n
	 稳定
*/

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
		// 选择子节点中较大的一个
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

// ShellSort 希尔排序
func ShellSort(list []int) {
	n := len(list)
	for d := n / 2; d > 0; d = d / 2 {
		for i := d; i < n; i++ {
			for j := i - d; j >= 0 && list[j+d] < list[j]; j = j - d {
				list[j], list[j+d] = list[j+d], list[j]
			}
		}
	}
}
