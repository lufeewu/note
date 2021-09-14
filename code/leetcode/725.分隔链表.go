/*
 * @lc app=leetcode.cn id=725 lang=golang
 *
 * [725] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	// 计算长度
	cur, n := root, 0
	for cur != nil {
		cur = cur.Next
		n++
	}

	var res []*ListNode
	size, extra := n/k, n%k
	cur = root
	for j := 0; j < k; j++ { // 注意点 1: 需要用 nil 补齐 k 个子链表
		res = append(res, cur)
		// 寻找下一个分隔点
		len := size
		if extra > 0 { // 注意点 2: 需要考虑链表长度不足 k 的整数倍情况, 所有链表长度差不超过 1
			len++
			extra--
		}
		for i := 0; i < len-1; i++ {
			cur = cur.Next
		}
		if cur != nil {
			cur.Next, cur = nil, cur.Next
		}
	}

	return res
}

// @lc code=end

