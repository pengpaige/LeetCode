/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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

// Iteration
// 后序遍历与前序遍历是逆序的，可以直接使用 144 题前序遍历结果的逆序
// 或者使用下面更复杂的方式直接后序遍历
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			stack = append(stack, top, nil)
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		} else if len(stack) > 0 {
			top = stack[len(stack)-1]
			ans = append(ans, top.Val)
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}

// Recursion
func postorderTraversal_Recursion(root *TreeNode) []int {
	var ans []int
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			f(node.Left)
		}
		if node.Right != nil {
			f(node.Right)
		}
		ans = append(ans, node.Val)
	}
	f(root)
	return ans
}

// @lc code=end

