/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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

// DFS
func isSubtree(s *TreeNode, t *TreeNode) bool {
	var ans bool
	var f func(s *TreeNode)
	f = func(s *TreeNode) {
        if s == nil {
            return
        }
		if s.Val == t.Val {
			if sameTree(s, t) {
				ans = true
				return
			}
		}
		f(s.Left)
		f(s.Right)
	}
	f(s)
	return ans
}

func sameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return sameTree(s.Left, t.Left) && sameTree(s.Right, t.Right)
}

// @lc code=end

