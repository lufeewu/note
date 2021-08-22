/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
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
func findFrequentTreeSum(root *TreeNode) []int {
	var res []int
	var middle = make(map[int]int)
	dfsFindFrequentTreeSum(root, middle)
	var maxNum = 0
	for _, v := range middle {
		if v > maxNum {
			maxNum = v
		}
	}
	for k, v := range middle {
		if v == maxNum {
			res = append(res, k)
		}
	}
	return res
}

func dfsFindFrequentTreeSum(root *TreeNode, middle map[int]int) int {
	if root == nil {
		return 0
	}
	left := dfsFindFrequentTreeSum(root.Left, middle)
	right := dfsFindFrequentTreeSum(root.Right, middle)
	cur := left + right + root.Val
	middle[cur]++
	return cur
}

// @lc code=end

