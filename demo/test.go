package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

func selectCase() {
	ch := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		ch[i] = make(chan int)
	}
	go func() {
		i := 0
		for {
			logrus.Infof("i=%d", i)
			time.Sleep(1 * time.Second)
			ch[i] <- 1
			i = (i + 1) % 3
		}
	}()

	for {
		select {
		case <-ch[0]:
			logrus.Infoln("0")
		case <-ch[1]:
			logrus.Infoln("1")
		case <-ch[2]:
			logrus.Infoln("2")
		default:
			logrus.Infoln("default")
			time.Sleep(1 * time.Second)
		}
	}
}

func testSize() {
	// selectCase()
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	var a []int
	var b map[int]int
	var c struct{}
	fmt.Println(unsafe.Sizeof(a)) // 24
	fmt.Println(unsafe.Sizeof(b)) // 8
	fmt.Println(unsafe.Sizeof(c)) // 0
	var ac uintptr
	fmt.Println(unsafe.Sizeof(ac)) // 8
}

// Time chan test
func Time() {
	type test struct {
		Cancel chan struct{}
	}
	t := test{
		Cancel: make(chan struct{}),
	}
	go func() {
		<-t.Cancel
		logrus.Infoln("1")
	}()
	time.Sleep(1 * time.Second)
	t.Cancel <- struct{}{}
	time.Sleep(1 * time.Second)
}

// GO 语言核心36讲-39 bytes 包的泄漏
func unreadBytesTest() {
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("%q %d\n", contents, buffer1.Cap())
	unreadBytes := buffer1.Bytes()
	fmt.Printf("%v \n", unreadBytes)

	buffer1.WriteString("cdefg")
	fmt.Printf("The Capacity of buffer:%d %v\n", buffer1.Cap(), string(buffer1.Bytes()))
	unreadBytes = unreadBytes[:cap(unreadBytes)] // leak new string from unreadBytes
	fmt.Printf("%v\n", string(unreadBytes))
}

func letsEncrypt() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	logrus.Infoln(autotls.Run(r, "127.0.0.1"))

}

func testClose() {

	cancel := make(chan struct{})
	cancel2 := make(chan struct{})

	go func() {
		// close(cancel2)
		time.Sleep(1 * time.Second)
		cancel <- struct{}{}
	}()

	select {
	case _, ok := <-cancel:
		logrus.Infoln("cancel1", ok)
	case _, ok := <-cancel2:
		logrus.Infoln("cancel2", ok)
	}

}

func testGoroutine() {
	wg := sync.WaitGroup{}
	var m int = 0
	// wg.Add(1000000) // 1 使用这行 TotalAlloc = 511 MiB 花费3s
	for i := 0; i < 1000000; i++ {
		wg.Add(1) //2  使用这行 TotalAlloc = 74 MiB 花费0.5s
		go func() {
			time.Sleep(time.Duration(1) * time.Millisecond)
			wg.Done()
			tmp := runtime.NumGoroutine()
			if m < tmp {
				m = tmp
			}
		}()
	}
	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem := mem.TotalAlloc / 1024 / 1024
	fmt.Printf("\tTotalAlloc = %v MiB  max goroutine: %d\n", curMem, m)
}
func testSlice() {
	// a := []int{1, 3, 4, 5, 6}
	// a = append(a, 4)
	// fmt.Println(a, cap(a), len(a))
	// b := make([]int, 2, 2)
	// b := a[2:4]
	// b = append(b, 5, 5, 5, 5, 5, 5, 5)
	// b[0] = 1
	// fmt.Println(b, cap(b), len(b))
	// fmt.Println(a, cap(a), len(a))
	var wg sync.WaitGroup
	wg.Add(12)
	for i := 0; i < 6; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	for i := 0; i < 6; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func testArray() {
	/*
			a := []int{1, 2, 3}
			b := make([]int, 0, 3)
			fmt.Println(b, len(b), cap(b))
			copy(a, b)
			fmt.Println(b)
		fmt.Println(selectPresent([][]int{
			[]int{1, 2, 3},
			[]int{2, 3, 4},
		}))
	*/
	fmt.Println(change("1234"))
}
func change(number string) string {
	// write code here
	result := []byte(number)
	i, j := 0, len(result)-1
	for i < j {
		for i < j && int(result[i]-'0')%2 != 0 {
			i++
		}
		for i < j && int(result[j]-'0')%2 != 0 {
			j--
		}
		if i < j {
			result[i], result[j] = result[j], result[i]
			i++
			j--
		}
	}
	return string(result)
}
func selectPresent(presentVolumn [][]int) int {
	// write code here
	if presentVolumn == nil {
		return 0
	}

	n, m := len(presentVolumn), len(presentVolumn[0])
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				if j == 0 {
					dp[i][j] = presentVolumn[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + presentVolumn[i][j]
				}
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + presentVolumn[i][j]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + presentVolumn[i][j]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1][m-1]
}

func min(value int, args ...int) int {
	fmt.Println("::", value, args)
	for _, v := range args {
		if v < value {
			value = v
		}
	}
	return value
}

func testInterface() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(1 * time.Second) // 0 - 9
	}
	// time.Sleep(1 * time.Second)  // 全 10
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	return searchSum(nums, 0, sum, sum)
}

func searchSum(nums []int, i int, sum1 int, sum2 int) bool {
	if sum1 == 0 || sum2 == 0 {
		return true
	}

	if i >= len(nums) || sum1 < 0 || sum2 < 0 {
		return false
	}
	fmt.Println(i, sum1, sum2)
	return searchSum(nums, i+1, sum1-nums[i], sum2) || searchSum(nums, i+1, sum1, sum2-nums[i])

}
func canPartition2(nums []int) bool {
	s := 0
	for _, n := range nums {
		s += n
	}
	if s%2 != 0 {
		return false
	}
	sort.Ints(nums)
	var d int
	r := DFS(nums, len(nums)-1, 0, 0, s, &d)
	fmt.Println(d)
	return r
}
func DFS(nums []int, i int, sum1, sum2 int, s int, d *int) bool { // sum1 : the sum of the 1st set; sum2: the sum of the 2nd set.
	*d++
	fmt.Println(*d)
	if sum1 == s/2 || sum2 == s/2 {
		return true
	}
	if i < 0 || sum1 > s/2 || sum2 > s/2 {
		fmt.Println("fff", i, sum1, sum2)
		return false
	}
	fmt.Println(i, sum1, sum2, s, d, nums[i])
	return DFS(nums, i-1, sum1+nums[i], sum2, s, d) || DFS(nums, i-1, sum1, sum2+nums[i], s, d)
}

func race() int {
	num := 1000
	wait := make(chan struct{}, num)
	n := 0
	for i := 0; i < num; i++ {
		go func() {
			// 译注：注意下面这一行
			n++ // 一次访问: 读, 递增, 写
			wait <- struct{}{}
		}()
	}

	// 译注：注意下面这一行
	// n++ // 另一次冲突的访问
	for i := 0; i < num; i++ {
		<-wait
	}
	fmt.Println(n) // 输出：未指定
	return n
}
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	set := make(map[string]bool)
	set["()"] = true
	for i := 1; i < n; i++ {
		set2 := make(map[string]bool)
		for k, _ := range set {
			tmp := fmt.Sprintf("()%s", k)
			set2[tmp] = true
			tmp = fmt.Sprintf("%s()", k)
			set2[tmp] = true
			tmp = fmt.Sprintf("(%s)", k)
			set2[tmp] = true
		}
		set = set2
	}
	result := make([]string, 0, len(set))
	for k, _ := range set {
		result = append(result, k)
	}
	return result
}

var letter = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	n := int(digits[0] - '0')
	fmt.Println(n)
	if n < 2 || n > 9 {
		return nil
	}
	count := 1
	for _, v := range digits {
		if v == '9' {
			count = count * 4
		} else {
			count = count * 3
		}
	}
	res := make([]string, 0, count)
	subRes := letterCombinations(digits[1:])
	for _, v := range letter[n-2] {
		if subRes == nil {
			res = append(res, string(v))
		} else {
			for _, subValue := range subRes {
				res = append(res, fmt.Sprintf("%s%s", string(v), subValue))
			}
		}

	}
	return res

}

// [3 1 2] 7
// 1 1 1 1 1 1 1
// 2 2 2 1
// 3 3 1 // 1 3 3
// 1 2 3 1

func targetArray(arr []int, target int) [][]int {
	if arr == nil {
		return nil
	}
	res := targetArrayHelper(arr, 0, target)
	return res
}

func targetArrayHelper(arr []int, start int, target int) [][]int {
	if start >= len(arr) || target <= 0 {
		return nil
	}
	var res [][]int

	pre := []int{}

	// 递归
	for tmp := 0; tmp <= target; tmp = tmp + arr[start] {
		if tmp != 0 {
			pre = append(pre, arr[start])
		}

		// end
		if target-tmp < 0 {
			break
		}

		if target-tmp == 0 {
			if pre != nil {
				res = append(res, pre)
			}
			break
		}

		// 递归
		subRes := targetArrayHelper(arr, start+1, target-tmp)
		for _, v := range subRes {
			v = append(v, pre...)
			res = append(res, v)
		}
	}

	return res
}

// leetcode 295
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
func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	s1, s2 := -2, -1
	subStones := make([]int, 0, len(stones)-1)
	for _, v := range stones {
		if v > s2 {
			if s1 > 0 {
				subStones = append(subStones, s1)
			}
			s1 = s2
			s2 = v
		} else if v > s1 {
			if s1 > 0 {
				subStones = append(subStones, s1)
			}
			s1 = v
		} else {
			subStones = append(subStones, v)
		}
	}
	if s2-s1 != 0 {
		subStones = append(subStones, s2-s1)
	}
	fmt.Println(s2, s1, subStones)
	return lastStoneWeight(subStones)
}

func circle(n int) int {
	var p int
	for i := 2; i <= n; i++ {
		p = (p + 3) % i
	}
	return p + 1
}
func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1 // 返回下标
	}
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func maxSubLen(s string) int {
	var resLen int
	table := make(map[byte]int) //  记录字符上一次最新出现的位置, 从 1 开始
	subStart := 0               // 子串的开始下标
	for k, v := range s {
		if table[byte(v)] >= subStart { // v 已经出现过，更新子串开启的位置
			subStart = table[byte(v)] // 子串的下标从重复字符的下一个开始
		}
		table[byte(v)] = k + 1
		if resLen < (k - subStart + 1) {
			resLen = k - subStart + 1
		}
	}

	return resLen

}

// 1. 无序 int 数组 []int{}  1, 2 ,3, 4, 5
func firstMissedNum(nums []int) int {
	for i, _ := range nums {
		swap(nums, i)
	}
	fmt.Println(nums)
	for k, v := range nums {
		if k != v && k > 0 {
			return k
		}
	}

	return -1
}

// func swap(nums []int, i int) {
// 	if i < len(nums) && i >= 0 && nums[i] != i &&
// 		nums[i] < len(nums) && nums[i] >= 0 {
// 		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
// 		swap(nums, nums[i])
// 	}
// }

func swap(nums []int, i int) {
	// time.Sleep(200 * time.Millisecond)
	if i < len(nums) && i >= 0 && nums[i] != (i+1) &&
		nums[i] <= len(nums) && nums[i] > 0 && nums[i] != nums[nums[i]-1] {
		// fmt.Printf("%v %v %v/ ", i, nums[i], nums[nums[i]-1])
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		swap(nums, i)
	}
}

func findDisappearedNumbers(nums []int) []int {
	var res []int
	for i, _ := range nums {
		swap(nums, i)
	}
	fmt.Println(nums)
	for k, v := range nums {
		if k+1 != v {
			res = append(res, k+1)
		}
	}

	return res
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{39, 8, 43, 12, 38, 11,
		-9, 12, 34, 20, 44, 32, 10, 22, 38, 9, 45, 26, -4, 2, 1, 3,
		3, 20, 38, 17, 20, 25, 41, 35, 37, 18, 37, 34, 24, 29, 39,
		9, 36, 28, 23, 18, -2, 28, 34, 30}))
	return
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{67, 1, -1, 3, -4, 6, 0}))
	fmt.Println(findDisappearedNumbers([]int{67, -1, 3, -4, 0, 6}))
	// fmt.Println(maxSubLen("dvdf"))
	// fmt.Println(binarySearch([]int{}, 8))

	return
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1, 5, 9}))
	return
	a := Constructor()
	a.AddNum(1)
	a.AddNum(2)
	fmt.Println(a.FindMedian())
	a.AddNum(3)
	fmt.Println(a.MaxHeap, a.MinHeap)
	fmt.Println(a.FindMedian())
	fmt.Println(a.FindMedian())
	a.AddNum(4)
	fmt.Println(a.FindMedian())
	a.AddNum(5)
	fmt.Println(a.FindMedian())
	a.AddNum(10)
	fmt.Println(a.FindMedian())
	a.AddNum(7)
	fmt.Println(a.FindMedian())
	a.AddNum(17)
	fmt.Println(a.FindMedian())
	return
	fmt.Println(targetArray([]int{3, 1, 2}, 4))
	return
	fmt.Println(letterCombinations("2"))
	return
	n := race()
	m := make(map[string]bool)
	m["sdf"] = true
	m["sd2f"] = true
	m["sd3f"] = true

	fmt.Println(len(m))
	result := make([]string, n)
	result = append(result)
	fmt.Println(result)
	// fmt.Println(canPartition2([]int{100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	// testSlice()
	// testGoroutine()
	// testArray()
	// testInterface()
	// t := Teacher{}
	// t.ShowA()
	// t.ShowB()

	return
	// Time()
	// t := time.NewTimer(10 * time.Second)
	// for {
	// 	t.Reset(1 * time.Second)
	// 	logrus.Infoln("1")
	// 	select {
	// 	case <-t.C:
	// 		logrus.Infoln("2")
	// 	}
	// }
	// unreadBytesTest()

	// letsEncrypt()
	// router := gin.Default()

	// router.RunTLS(":18080", "", "")
	// testClose()

	r := gin.New()
	r.GET("/a/b/c", func(c *gin.Context) {
		s := strings.Split(c.Request.URL.String(), "/")
		fmt.Println("t:", s[len(s)-1], len(s))
		c.JSON(200, nil)
	})
	err := r.Run("0.0.0.0:12222")
	fmt.Println(err)

}
