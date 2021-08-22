/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
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

type item struct {
	idx int // 记录下标
	*TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans, que := 1, []item{{0, root}}
	for len(que) > 0 {
		if l := que[len(que)-1].idx - que[0].idx + 1; l > ans {
			ans = l
		}
		tmp := []item{}
		for _, q := range que {
			if q.Left != nil {
				tmp = append(tmp, item{q.idx * 2, q.Left})
			}
			if q.Right != nil {
				tmp = append(tmp, item{q.idx*2 + 1, q.Right})
			}
		}
		que = tmp
	}
	return ans
}

// @lc code=end

