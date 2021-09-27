/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	var nums []int
	dfsHelper(root, 0, &nums)
	var res = 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}

func dfsHelper(root *TreeNode, value int, nums *[]int) {
	if root == nil {
		return
	}
	value = value*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*nums = append(*nums, value)
		return
	}
	dfsHelper(root.Left, value, nums)
	dfsHelper(root.Right, value, nums)
}

// @lc code=end

