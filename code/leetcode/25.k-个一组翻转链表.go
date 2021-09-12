/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	var cur, next *ListNode = head, head.Next

	cur.Next = nil
	i := k - 1
	for i > 0 && cur != nil && next != nil {
		i--
		cur, next, next.Next = next, next.Next, cur
	}

	if next != nil {
		head.Next = reverseKGroup(next, k)
	} else if i > 0 {
		return reverseKGroup(cur, k-i)
	}

	return cur
}

// @lc code=end

