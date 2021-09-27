/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	dfs(root, &res)
	return res - 1
}

func dfs(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, max)
	r := dfs(root.Right, max)
	if l+r+1 > *max {
		*max = l + r + 1
	}
	if l < r {
		return r + 1
	}
	return l + 1
}

// @lc code=end

