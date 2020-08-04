/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
// 二叉搜索树通过中序遍历可以对节点的值进行从小到大的排序
func kthSmallest(root *TreeNode, k int) int {
	vlist := make([]int, 0)
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil || len(vlist) >= k {
			return
		}
		inOrder(node.Left)
		vlist = append(vlist, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return vlist[k-1]
}

// @lc code=end

