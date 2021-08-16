/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type IntMaxHeap []int

func (h IntMaxHeap) Len() int {
	return len(h)
}

func (h IntMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	MaxHeap IntMaxHeap
	MinHeap IntMinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{}
	return m
}

func (this *MedianFinder) AddNum(num int) {

	// len max 1
	if this.MaxHeap.Len() > this.MinHeap.Len() {
		heap.Push(&this.MinHeap, num)
	} else {
		heap.Push(&this.MaxHeap, num)
	}

	// adjust
	for this.MaxHeap.Len() > 0 && this.MinHeap.Len() > 0 &&
		this.MaxHeap[0] > this.MinHeap[0] {
		tmp := heap.Pop(&this.MaxHeap)
		heap.Push(&this.MaxHeap, heap.Pop(&this.MinHeap))
		heap.Push(&this.MinHeap, tmp)
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.MaxHeap.Len() == 0 {
		return 0
	}
	if this.MinHeap.Len() == 0 {
		return float64(this.MaxHeap[0])
	}
	if this.MaxHeap.Len() == this.MinHeap.Len() {
		return float64(this.MaxHeap[0]+this.MinHeap[0]) / 2
	} else {
		return float64(this.MaxHeap[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

