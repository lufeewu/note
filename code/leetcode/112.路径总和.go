/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumDFS(root, targetSum)
}
func hasPathSumDFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var stackNodes []*TreeNode
	var stackSum []int

	var node *TreeNode = root
	var sum int
	for node != nil || len(stackNodes) != 0 {
		for node != nil {
			sum += node.Val
			stackSum = append(stackSum, sum)
			stackNodes = append(stackNodes, node)
			node = node.Left
		}

		if len(stackNodes) != 0 {
			node = stackNodes[len(stackNodes)-1]
			stackNodes = stackNodes[:len(stackNodes)-1]

			sum = stackSum[len(stackSum)-1]
			stackSum = stackSum[:len(stackSum)-1]

			if node.Left == nil && node.Right == nil &&
				sum == targetSum {
				return true
			}

			node = node.Right
		}

	}

	return false
}
func hasPathSumBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var stackNode []*TreeNode
	var stackSum []int

	stackNode = append(stackNode, root)
	stackSum = append(stackSum, root.Val)

	var node *TreeNode
	var sum int
	for len(stackNode) != 0 {
		node = stackNode[len(stackNode)-1]
		stackNode = stackNode[:len(stackNode)-1]

		sum = stackSum[len(stackSum)-1]
		stackSum = stackSum[:len(stackSum)-1]

		if node.Left == nil && node.Right == nil &&
			sum == targetSum {
			return true
		}

		if node.Left != nil {
			stackNode = append(stackNode, node.Left)
			stackSum = append(stackSum, sum+node.Left.Val)
		}
		if node.Right != nil {
			stackNode = append(stackNode, node.Right)
			stackSum = append(stackSum, sum+node.Right.Val)
		}
	}

	return false
}

// @lc code=end

