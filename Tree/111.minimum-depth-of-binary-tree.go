/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if root.Left == nil && root.Right == nil{
		return 1
	}
	// 左子树为空或者右子树为空需要特殊处理
	// 否则直接按照最大深度的解法会导致认为有深度为 0 的子树
	// 而实际上不存在这种子树，比如这个 case [1, 2]
	if root.Left == nil {
		return r+1
	}
	if root.Right == nil {
		return l+1
	}
	return min(l, r) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

