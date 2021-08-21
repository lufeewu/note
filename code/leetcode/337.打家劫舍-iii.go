/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	v1, v2 := helper(root)
	return max(v1, v2)
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftValue, leftValue1 := helper(root.Left)
	rightValue, rightValue1 := helper(root.Right)
	value := max(leftValue, leftValue1) + max(rightValue, rightValue1) // 不偷盗根节点
	value1 := root.Val + leftValue + rightValue                        // 偷盗根节点
	return value, value1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

