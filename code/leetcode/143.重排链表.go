/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 寻找中间节点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 分隔链表
	fast = slow.Next
	slow.Next = nil

	// 反转第二段链表
	slow = nil
	for fast != nil {
		fast.Next, fast, slow = slow, fast.Next, fast
	}

	// 重排链表

	slow, fast = head, slow
	for fast != nil {
		slow.Next, slow, fast.Next, fast = fast, slow.Next,
			slow.Next, fast.Next
	}

}

// @lc code=end

