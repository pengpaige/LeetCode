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

//  Iteration
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			ans = append(ans, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr = top.Right
	}
	return ans
}

func preorderTraversal_Recursion(root *TreeNode) []int {
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

