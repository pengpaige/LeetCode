/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var findSum func(node *TreeNode, curr int) bool
	findSum = func(node *TreeNode, curr int) bool {
		// 从 root 开始每次判断左右子节点是否 nil 就不会遇到 node == nil
		if node.Left == nil && node.Right == nil {
			if curr+node.Val == sum {
				return true
			}
			return false
		}
		// 和求最小路径和一样，类似 [1, 2] 的 case 只有一条路径
		// 只有根节点不算一条完整的路径
		if node.Left == nil {
			return findSum(node.Right, curr+node.Val)
		}
		if node.Right == nil {
			return findSum(node.Left, curr+node.Val)
		}
		return findSum(node.Left, curr+node.Val) || findSum(node.Right, curr+node.Val)
	}
	return findSum(root, 0)
}

// @lc code=end

