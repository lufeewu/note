/*
 * @lc app=leetcode.cn id=1019 lang=golang
 *
 * [1019] 链表中的下一个更大节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	// 维护一个单调递减栈, 知道找到递增的数据，则填充栈顶更小元素的下一个最大节点
	stack, ans := make([]int, len(data)), make([]int, len(data))

	for index := 0; index < len(data); index++ {
		if len(stack) == 0 {
			stack = append(stack, index)
		} else {
			for len(stack) > 0 && data[index] > data[stack[len(stack)-1]] {
				pop := len(stack) - 1
				ans[stack[pop]] = data[index]
				stack = stack[:pop]
			}
			stack = append(stack, index)
		}
	}
	return ans
}

// @lc code=end

