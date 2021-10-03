/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur = head, head.Next
	for cur != nil {
		if cur.Val == pre.Val {
			pre.Next = cur.Next
			cur.Next = nil
			cur = pre.Next
		} else {
			pre, cur = cur, cur.Next
		}
	}
	return head
}

// @lc code=end

