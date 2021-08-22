/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	value, _ := findBottomLeftValueHelper(root)
	return value
}

func findBottomLeftValueHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val, 1
	}
	leftValue, leftHeight := findBottomLeftValueHelper(root.Left)
	rightValue, rightHeight := findBottomLeftValueHelper(root.Right)
	if leftHeight >= rightHeight {
		return leftValue, leftHeight + 1
	}
	return rightValue, rightHeight + 1

}

// @lc code=end

