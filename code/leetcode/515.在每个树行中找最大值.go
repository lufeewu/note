/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	// 结束条件
	if root == nil {
		return nil
	}
	//迭代
	result := make([]int, 0)
	//通过queue对层次数据的维护
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		queueLength := len(queue)
		max := math.MinInt64
		for i := 0; i < queueLength; i++ {
			if queue[i].Val > max {
				max = queue[i].Val
			}
			// 在遍历本层的基础上，同时把下一层的节点记录到queue中
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		// 把每一层的最大值，添加到result
		result = append(result, max)
		// 更新queue，保证数据为下一层所需的节点数据
		queue = queue[queueLength:]
	}
	return result
}

// @lc code=end

