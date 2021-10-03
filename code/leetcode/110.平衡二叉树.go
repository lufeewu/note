/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	r, rig := isBalancedHelper(root.Right)
	l, lef := isBalancedHelper(root.Left)
	if r > l+1 || l > r+1 {
		return false
	}
	x
	return rig && lef
}

func isBalancedHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	lh, lb := isBalancedHelper(root.Left)
	rh, rb := isBalancedHelper(root.Right)
	if rh > lh+1 || lh > rh+1 {
		return 0, false
	}
	if lh >= rh {
		return 1 + lh, lb && rb
	} else {
		return 1 + rh, lb && rb
	}

}

// @lc code=end

