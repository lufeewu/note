/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	var min []int
	helper(root, &min)
	return min[k-1]
}

func helper(root *TreeNode, min *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, min)
	*min = append(*min, root.Val)
	helper(root.Right, min)
}

// @lc code=end

