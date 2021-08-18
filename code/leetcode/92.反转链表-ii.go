/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var begin *ListNode = head
	var pre *ListNode = nil
	i := 1
	for head != nil {
		if i < m {
			pre, head = head, head.Next
			i++
		} else {
			var rPre *ListNode = pre
			for i <= n {
				rPre, head, head.Next = head, head.Next, rPre
				i++
			}
			if m == 1 && head == nil {
				return rPre
			}
			if pre != nil {
				pre.Next.Next = head
				pre.Next = rPre
				return begin
			} else {
				begin.Next = head
				return rPre
			}
		}
	}
	return begin
}

// @lc code=end

