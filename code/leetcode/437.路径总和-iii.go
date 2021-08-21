/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, sum int) int {
	var res = 0
	var nodeArray []*TreeNode
	helper(root, sum, 0, nodeArray, &res)
	return res
}

func helper(node *TreeNode, sum int, curSum int, nodeArray []*TreeNode, res *int) {
	if node == nil {
		return
	}
	nodeArray = append(nodeArray, node)
	curSum = curSum + node.Val
	if curSum == sum {
		*res = *res + 1
	}

	var t = curSum
	for i := 0; i < len(nodeArray)-1; i++ {
		t = t - nodeArray[i].Val
		if t == sum {
			*res = *res + 1
		}
	}

	helper(node.Left, sum, curSum, nodeArray, res)
	helper(node.Right, sum, curSum, nodeArray, res)
	nodeArray = nodeArray[0 : len(nodeArray)-1]
}

// @lc code=end

