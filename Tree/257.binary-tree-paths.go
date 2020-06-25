/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	var dfs func(node *TreeNode, curr string)
	dfs = func(node *TreeNode, curr string) {
		// root == nil 或者叶子节点的空孩子
		if node == nil {
			return
		}
		if node == root {
			curr += strconv.Itoa(node.Val)
		} else {
			curr += "->" + strconv.Itoa(node.Val)
		}
		// 当前这一层就要结束防止将左或右子树之一为空误判为叶子节点
		if node.Left == nil && node.Right == nil {
			ans = append(ans, curr)
		}
		dfs(node.Left, curr)
		dfs(node.Right, curr)
	}
	dfs(root, "")
	return ans
}
// @lc code=end

