/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		n := queue.Len()
		var tmp []int
		for i := 0; i < n; i++ {
			e := queue.Front()
			if e != nil {
				queue.Remove(e)
				eNode, ok := e.Value.(*TreeNode)
				if !ok || eNode == nil {
					continue
				}
				tmp = append(tmp, eNode.Val)
				queue.PushBack(eNode.Left)
				queue.PushBack(eNode.Right)
			}
		}
		if tmp != nil {
			res = append(res, tmp)
		}
	}
	return res
}

// @lc code=end

