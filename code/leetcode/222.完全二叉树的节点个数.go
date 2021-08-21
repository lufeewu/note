/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int { // 返回以root为根节点的子树的节点个数
	if root == nil { // 递归的出口
		return 0
	}
	lH, rH := 0, 0             // 两侧高度
	lNode, rNode := root, root // 两个指针

	for lNode != nil { // 计算左侧高度
		lH++
		lNode = lNode.Left
	}
	for rNode != nil { // 计算右侧高度
		rH++
		rNode = rNode.Right
	}
	if lH == rH { // 当前子树是满二叉树，返回出节点数
		return 1<<lH - 1 // 左移n位就是乘以2的n次方
	}
	// 当前子树不是完美二叉树，只是完全二叉树，递归处理左右子树
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// @lc code=end

