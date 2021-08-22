/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] < nums[i] {
			index = i
		}
	}
	root := &TreeNode{Val: nums[index]}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

// @lc code=end

