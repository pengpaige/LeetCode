/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var level []*TreeNode
	level = append(level, root)
	for len(level) > 0 {
		var tmpLevel []*TreeNode
		var tmpVals []int
		for _, node := range level {
			if node == nil {
				continue
			}
			tmpVals = append(tmpVals, node.Val)
			tmpLevel = append(tmpLevel, node.Left, node.Right)
		}
		if tmpVals != nil {
			// 往前插
			ans = append([][]int{tmpVals}, ans...)
		}
		level = tmpLevel
	}
	return ans
}

// @lc code=end

