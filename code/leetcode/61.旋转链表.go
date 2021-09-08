/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var n int = 0
	var cur, end *ListNode = head, nil
	for cur != nil {
		n++
		if cur.Next == nil {
			end = cur
		}
		cur = cur.Next
	}

	k = k % n
	if k == 0 {
		return head
	}

	cur = head
	for i := 0; i < n-k-1; i++ {
		cur = cur.Next
	}

	end.Next = head
	head = cur.Next
	cur.Next = nil

	return head
}

// @lc code=end

