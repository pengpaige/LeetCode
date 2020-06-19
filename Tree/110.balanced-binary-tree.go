/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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

// Bottom Up, time 81%
// 思路和 104 直接求最大深度差不多，只是要在求深度的同时判断是否平衡
func isBalanced(root *TreeNode) bool {
    var f func(node *TreeNode) (bool, int)
	f = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		leftIsBalanced, leftDepth := f(node.Left)
		if !leftIsBalanced {
			return false, 0
		}
		rightIsBalanced, rightDepth := f(node.Right)
		if !rightIsBalanced {
			return false, 0
		}
		if !checkABS(leftDepth, rightDepth) {
			return false, 0
		}
		return true, max(leftDepth, rightDepth)+1
	}
    balanced, _ := f(root)
    return balanced
}

// time 19%
func isBalanced_TopDown(root *TreeNode) bool {
	mem := make(map[*TreeNode]bool, 0)
	var f func(node *TreeNode) bool
	f = func(node *TreeNode) bool {
		if b, ok := mem[node]; ok {
			return b
		}
		if node == nil {
			return true
		}
		if !checkABS(maxDepth(node.Left), maxDepth(node.Right)) {
			return false
		}
        ret := f(node.Left) && f(node.Right)
		mem[node] = ret
		return ret
	}
	return f(root)
}

func checkABS(d1, d2 int) bool {
	diff := d1 - d2
	if diff <= 1 && diff >= -1 {
		return true
	}
	return false
}

// problem 104
func maxDepth(root *TreeNode) int {
	mem := make(map[*TreeNode]int, 0)
	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
		if depth, ok := mem[node]; ok {
			return depth
		}
		if node == nil {
			return 0
		}
		d := max(f(node.Left), f(node.Right)) + 1
		mem[node] = d
		return d
	}
	return f(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

