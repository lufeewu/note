/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	var curValue, carry = 0, 0

	var l1Cur, l2Cur = l1, l2
	var head, cur *ListNode
	for l1Cur != nil || l2Cur != nil {
		curValue = 0
		if l1Cur != nil {
			curValue += l1Cur.Val
			l1Cur = l1Cur.Next
		}
		if l2Cur != nil {
			curValue += l2Cur.Val
			l2Cur = l2Cur.Next
		}

		curValue = curValue + carry
		carry, curValue = curValue/10, curValue%10
		if cur == nil {
			cur = &ListNode{
				Val: curValue,
			}
			head = cur
		} else {
			cur.Next = &ListNode{
				Val: curValue,
			}
			cur = cur.Next
		}

	}
	if carry != 0 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}
	head = reverseList(head)
	return head
}

func reverseList(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	var pre, next *ListNode
	for l != nil {
		next = l.Next
		l.Next = pre
		pre = l
		l = next
	}
	l = pre
	return pre
}

// @lc code=end

