/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	var cur, tmp, res *Node = head, nil, nil

	for cur != nil {
		tmp = &Node{
			Val:    cur.Val,
			Next:   cur.Next,
			Random: cur.Random,
		}
		cur.Next, cur = tmp, tmp.Next
	}

	cur = head.Next
	for cur != nil {
		if cur.Random != nil {
			cur.Random = cur.Random.Next
		}
		if cur.Next == nil {
			break
		}
		cur = cur.Next.Next
	}

	cur = head
	res = cur.Next
	for cur != nil && cur.Next != nil {
		if cur.Next.Next != nil {
			cur.Next, cur.Next.Next = cur.Next.Next, cur.Next.Next.Next
			cur = cur.Next
		} else {
			cur.Next = nil
		}
	}
	return res

}

// @lc code=end

