/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	return sortedListToBSTHelper(head, n)

}

func sortedListToBSTHelper(head *ListNode, n int) *TreeNode {
	if n == 0 {
		return nil
	}
	leftLen, rightLen := n/2, n-1-n/2
	rootNode := head
	for i := 0; i < leftLen; i++ {
		rootNode = rootNode.Next
	}

	var root = &TreeNode{
		Val: rootNode.Val,
	}
	root.Left = sortedListToBSTHelper(head, leftLen)
	root.Right = sortedListToBSTHelper(rootNode.Next, rightLen)
	return root
}

// @lc code=end

