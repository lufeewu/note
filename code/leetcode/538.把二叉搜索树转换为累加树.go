/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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

func convertBST(root *TreeNode) *TreeNode {
	var sum = 0
	convertBFTHelper(root, &sum)
	return root
}

func convertBFTHelper(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	convertBFTHelper(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	convertBFTHelper(root.Left, sum)
	return
}

// @lc code=end

