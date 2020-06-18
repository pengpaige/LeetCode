/*
 * @lc app=leetcode id=1302 lang=golang
 *
 * [1302] Deepest Leaves Sum
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
func deepestLeavesSum(root *TreeNode) int {
	return dls(root, maxDepth(root))
}

func dls(root *TreeNode, maxDepth int) int {
	var ans int
	var f func(node *TreeNode, depth int)
	f = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Left != nil {
			f(node.Left, depth+1)
		}
		if node.Right != nil {
			f(node.Right, depth+1)
		}

		if node.Left == nil && node.Right == nil && depth == maxDepth {
			ans += node.Val
		}
	}
	f(root, 1)
	return ans
}

// post traverse
func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var left, right int
	if node.Left != nil {
		left = maxDepth(node.Left)
	}
	if node.Right != nil {
		right = maxDepth(node.Right)
	}
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

