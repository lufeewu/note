/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	var pre, cur *ListNode = nil, head
	for i := 0; i < n; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	if cur == nil {
		head.Next, head = nil, head.Next
		return head
	}
	pre = head
	cur = cur.Next
	for cur != nil {
		pre, cur = pre.Next, cur.Next
	}
	pre.Next, pre.Next.Next = pre.Next.Next, nil
	return head
}

// @lc code=end

