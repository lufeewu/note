/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	maxPath := -(1 << 31)
	dfsMaxSum(root, &maxPath)
	return maxPath
}

func dfsMaxSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	l := max(0, dfsMaxSum(root.Left, maxSum))
	r := max(0, dfsMaxSum(root.Right, maxSum))
	*maxSum = max(*maxSum, l+r+root.Val)
	return root.Val + max(l, r)

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

