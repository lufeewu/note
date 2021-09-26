/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}
	i := 0
	result := &TreeNode{
		Val: preorder[0],
	}
	for i < len(inorder) && inorder[i] != result.Val {
		i++
	}
	if i < len(inorder) {
		result.Left = buildTree(preorder[1:i+1], inorder[0:i])
		result.Right = buildTree(preorder[i+1:len(preorder)], inorder[i+1:len(inorder)])
	}

	return result
}

// @lc code=end

