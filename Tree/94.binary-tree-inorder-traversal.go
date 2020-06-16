/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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

// iteration
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	var stack []*TreeNode
	stack = append(stack, root)
	curr := root
	for curr != nil && len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		ans = append(ans, curr.Val)
		stack = stack[:len(stack)-1]

		if curr.Right != nil {
			curr = curr.Right
		}
	}
	return ans
}

func inorderTraversal_Recursion(root *TreeNode) []int {
	var ans []int
	var f func(root *TreeNode)
	f = func(node *TreeNode) {
		// 这里只对根节点有效，其他节点直接在判断左右时拦截
		if node == nil {
			return
		}
		if node.Left != nil {
			f(node.Left)
		}
		ans = append(ans, node.Val)
		if node.Right != nil {
			f(node.Right)
		}
		return
	}
	f(root)
	return ans
}

// @lc code=end

