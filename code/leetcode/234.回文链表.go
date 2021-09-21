/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	n := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}
	cur = head
	for i := 0; i < n/2; i++ {
		cur = cur.Next
	}

	var pre *ListNode = nil
	for cur != nil {
		cur.Next, cur, pre = pre, cur.Next, cur
	}

	cur = pre
	for i := 0; i < n/2; i++ {
		if cur.Val != head.Val {
			return false
		}
		cur = cur.Next
		head = head.Next
	}

	return true
}

// @lc code=end

