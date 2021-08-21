/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	oddHead := &ListNode{
		Val: 0,
	}
	evenHead := &ListNode{
		Val: 0,
	}

	odd, even := oddHead, evenHead

	for head != nil {
		odd.Next = head
		even.Next = head.Next

		odd = odd.Next
		even = even.Next

		if head.Next != nil {
			head = head.Next.Next
		} else {
			head = head.Next
		}

	}

	odd.Next = evenHead.Next
	return oddHead.Next
}

// @lc code=end

