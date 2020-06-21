/*
 * @lc app=leetcode id=669 lang=golang
 *
 * [669] Trim a Binary Search Tree
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
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	return doTrimBST(root, L, R)
}

func doTrimBST(node *TreeNode, L, R int) *TreeNode {
	if node == nil {
		return nil
	}
	for node != nil && !inLR(node.Val, L, R) {
		if node.Val < L {
			node = node.Right
		} else if node.Val > R { // 这里必须要二选一否则可能出现 空指针引用 比如 case: [1, 0, 2]
			node = node.Left
		}
	}
	if node == nil {
		return nil
	}
	node.Left = doTrimBST(node.Left, L, R)
	node.Right = doTrimBST(node.Right, L, R)
	return node
}

func inLR(x, l, r int) bool {
	if x >= l && x <= r {
		return true
	}
	return false
}
// @lc code=end

