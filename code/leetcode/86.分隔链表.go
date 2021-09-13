/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var p1Head, p1, p2Head, p2 *ListNode
	for head != nil {
		if head.Val < x {
			if p1Head == nil {
				p1Head = head
				p1 = p1Head
			} else {
				p1.Next = head
				p1 = head
			}
		} else {
			if p2Head == nil {
				p2Head = head
				p2 = p2Head
			} else {
				p2.Next = head
				p2 = head
			}
		}
		head = head.Next
	}
	if p1Head == nil {
		return p2Head
	}
	if p2Head == nil {
		return p1Head
	}
	p2.Next = nil
	p1.Next = p2Head
	return p1Head
}

// @lc code=end

