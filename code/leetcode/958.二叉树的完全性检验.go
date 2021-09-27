/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] 二叉树的完全性检验
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
func isCompleteTree(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}
	var isBottom = false
	var stack []*TreeNode
	stack = append(stack, root)
	var n = len(stack)
	for n != 0 {
		for i := 0; i < n; i++ {
			if isBottom == true && (stack[i].Left != nil || stack[i].Right != nil) {
				return false
			} else if isBottom == true {
				continue
			}

			// 叶子节点非最后一层
			if stack[i].Left == nil && stack[i].Right == nil {
				isBottom = true
			} else if stack[i].Left != nil && stack[i].Right != nil {
				stack = append(stack, stack[i].Left)
				stack = append(stack, stack[i].Right)
			} else if stack[i].Right != nil {
				return false
			} else {
				isBottom = true
				stack = append(stack, stack[i].Left)
			}
		}
		stack = stack[n:]
		n = len(stack)
	}

	return true
}

// @lc code=end

