/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		if node.Left != nil {
			f(node.Left)
		}
		if node.Right != nil {
			f(node.Right)
		}
		return
	}
	f(root)
	return ans
}
// @lc code=end

