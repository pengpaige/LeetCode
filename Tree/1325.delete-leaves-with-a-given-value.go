/*
 * @lc app=leetcode id=1325 lang=golang
 *
 * [1325] Delete Leaves With a Given Value
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
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	// 剪掉原来的叶子节点之后新生成的叶子节点满足要求也要删除
	// 这个要求暗示了必须要用后序遍历
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		root = nil
	}
	return root
}

// @lc code=end

