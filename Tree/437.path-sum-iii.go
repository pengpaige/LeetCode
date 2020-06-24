/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func dfs(node *TreeNode, curr int) int {
	if node == nil {
		return 0
	}
	childCount := dfs(node.Left, curr-node.Val) + dfs(node.Right, curr-node.Val)
	if curr-node.Val == 0 {
		return 1 + childCount
	}
	return childCount
}

// @lc code=end

