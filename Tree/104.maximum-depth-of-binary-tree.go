/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	mem := make(map[*TreeNode]int, 0)
	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
		if depth, ok := mem[node]; ok {
			return depth
		}
		if node == nil {
			return 0
		}
		d := max(f(node.Left), f(node.Right)) + 1
		mem[node] = d
		return d
	}
	return f(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

