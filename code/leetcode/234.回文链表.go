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
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow.Next
	slow = nil
	for fast != nil {
		fast.Next, fast, slow = slow, fast.Next, fast
	}

	slow, fast = head, slow
	for fast != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow, fast = slow.Next, fast.Next
	}

	return true
}

// @lc code=end

