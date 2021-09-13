/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
	var res = &ListNode{}
	var cur = res
	var pre *ListNode
	pre = head
	for head != nil {
		head = head.Next
		if head != nil && head.Val == pre.Val {
			for head != nil && head.Val == pre.Val {
				head = head.Next
			}
			pre = head
		} else {
			pre.Next = nil
			cur.Next = pre
			cur = pre
			pre = head
		}

	}
	return res.Next
}

// @lc code=end

