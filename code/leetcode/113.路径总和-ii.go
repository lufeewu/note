/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}

	}
	var routes []int
	return pathSumHelperDFS(root, targetSum, routes)
}

func pathSumHelperDFS(root *TreeNode, targetSum int, routes []int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	routes = append(routes, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		var r = make([]int, len(routes))
		copy(r, routes)
		res = append(res, r)
		return res
	}

	left := pathSumHelperDFS(root.Left, targetSum-root.Val, routes)
	res = append(res, left...)
	right := pathSumHelperDFS(root.Right, targetSum-root.Val, routes)
	res = append(res, right...)

	return res
}

// @lc code=end

