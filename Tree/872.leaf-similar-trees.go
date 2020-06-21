/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1, leafs2 := getLeafList(root1), getLeafList(root2)
	if len(leafs1) != len(leafs2) {
		return false
	}
	for i := range leafs1 {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}
	return true
}

func getLeafList(root *TreeNode) []int {
	var ans []int
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			ans = append(ans, node.Val)
		}
		f(node.Left)
		f(node.Right)
	}
	f(root)
	return ans
}
// @lc code=end

