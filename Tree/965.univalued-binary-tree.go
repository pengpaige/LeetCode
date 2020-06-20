/*
 * @lc app=leetcode id=965 lang=golang
 *
 * [965] Univalued Binary Tree
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
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
		return false
	}
	uniVal := root.Val
	var f func(node *TreeNode) bool 
	f = func(node *TreeNode) bool {
		// root == nil 已经在上面被排除调了
		// 这里在出现 nil 肯定是遍历完了一个分支，直接返回真
		if node == nil {
			return true
		}
		if node.Val != uniVal {
			return false
		}
		return f(node.Left) && f(node.Right)
	}
	return f(root)
}
// @lc code=end

