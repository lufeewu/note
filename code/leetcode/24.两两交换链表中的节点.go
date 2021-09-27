/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, one, two *ListNode = nil, head, head.Next
	head = two
	for one != nil && one.Next != nil {
		two = one.Next
		one.Next, two.Next = two.Next, one
		if pre != nil {
			pre.Next = two
		}
		pre = one
		one = one.Next
	}
	return head
}

// @lc code=end

