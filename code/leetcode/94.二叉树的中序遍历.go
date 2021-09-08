/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stackNodes []*TreeNode
	var node *TreeNode = root
	for node != nil || len(stackNodes) != 0 {
		for node != nil {
			stackNodes = append(stackNodes, node)
			node = node.Left
		}
		if len(stackNodes) != 0 {
			node = stackNodes[len(stackNodes)-1]
			stackNodes = stackNodes[:len(stackNodes)-1]

			res = append(res, node.Val)
			node = node.Right
		}

	}

	return res
}

// 递归
func inorderTraversalRecursion(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// @lc code=end

