package sort

/*
 1. 简单排序
	- 选择排序
	- 插入排序
	- 冒泡排序
*/
// 可参考文章: https://www.cnblogs.com/gaopeng527/p/6699648.html

// 选择排序
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
func merge(list []int, a, b int) {
	var (
		i, j int
		mid  int
		pos  int
	)

	temp := make([]int, len(list))
	copy(temp, list)
	i = a
	mid = (a + b) / 2
	j = mid + 1
	pos = a
	for i <= mid && j <= b {
		if temp[i] < temp[j] {
			list[pos] = temp[i]
			pos++
			i++
		} else {
			list[pos] = temp[j]
			pos++
			j++
		}
	}
	for i <= mid {
		list[pos] = temp[i]
		pos++
		i++
	}
	for j <= b {
		list[pos] = temp[j]
		pos++
		j++
	}
}

// 左闭右闭[a,b]
func MergeSort(list []int, a, b int) {
	if a+1 >= b {
		if list[a] > list[b] {
			list[a], list[b] = list[b], list[a]
		}
		return
	}
	// 排序包括下标(a+b)/2的值
	MergeSort(list, a, (a+b)/2)

	MergeSort(list, (a+b)/2+1, b)

	merge(list, a, b)
}

/*
  3. 快速排序
*/
func QuickSort(list []int) {
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

		QuickSort(list[0 : head-1])
		QuickSort(list[head:])
	} else {
		list[0], list[head] = list[head], list[0]

		QuickSort(list[0:head])
		QuickSort(list[head+1:])
	}

}

/*
  4. 堆排序
	 nlog n
	 稳定
*/
func HeapSort(list []int) {
	n := len(list)
	// 建堆
	for i := n/2 - 1; i >= 0; i-- {
		k := i
		for 2*k+1 < n {
			j := 2*k + 1
			if j+1 < n && list[j] < list[j+1] {
				j++
			}
			if list[j] > list[k] {
				list[k], list[j] = list[j], list[k]
				k = j
			} else {
				break
			}
		}
	}

	// 调整堆
	for i := n - 1; i > 0; i-- {
		list[0], list[i] = list[i], list[0]
		k := 0
		for 2*k+1 < i {
			j := 2*k + 1
			if j+1 < i && list[j] < list[j+1] {
				j++
			}
			if list[j] > list[k] {
				list[k], list[j] = list[j], list[k]
				k = j
			} else {
				break
			}
		}
	}
}

/*
  5. 希尔排序
*/
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
