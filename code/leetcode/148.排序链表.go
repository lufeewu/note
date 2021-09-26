/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dump = &ListNode{
		Val:  -1,
		Next: head,
	}

	var pslow, pfast = dump, dump
	for pfast != nil && pfast.Next != nil {
		pslow = pslow.Next
		pfast = pfast.Next.Next
	}
	pfast = pslow.Next
	pslow.Next = nil
	var l1, l2 = sortList(head), sortList(pfast)
	return mergeSortLists(l1, l2)
}

func mergeSortLists(l *ListNode, r *ListNode) *ListNode {
	var dump = &ListNode{
		Val: -1,
	}
	var h1, h2, tail = l, r, dump
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			tail.Next = h1
			h1 = h1.Next
		} else {
			tail.Next = h2
			h2 = h2.Next
		}
		tail = tail.Next
	}

	if h1 == nil {
		tail.Next = h2
	} else if h2 == nil {
		tail.Next = h1
	}

	return dump.Next
}

// @lc code=end

