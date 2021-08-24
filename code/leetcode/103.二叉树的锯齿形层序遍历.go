/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	// init
	stack := []*TreeNode{}
	stack = append(stack, root)
	reverse := false

	// level order traversal
	for len(stack) != 0 {
		tmp := []*TreeNode{}
		levelRes := make([]int, 0, len(stack))
		for _, v := range stack {
			levelRes = append(levelRes, v.Val)
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}

		// reverse
		if reverse {
			reverseSlice(levelRes)
		}

		res = append(res, levelRes)

		reverse = !reverse
		stack = tmp
	}
	return res
}

func reverseSlice(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// @lc code=end

