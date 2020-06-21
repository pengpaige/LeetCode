/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var level []*TreeNode
	level = append(level, root)
	for len(level) > 0 {
		var tmpVals []int
		var tmpNodes []*TreeNode
		for i := range level {
			if level[i] == nil {
				continue
			}
			tmpVals = append(tmpVals, level[i].Val)
			tmpNodes = append(tmpNodes, level[i].Left, level[i].Right)
		}
		if tmpVals != nil {
			ans = append(ans, tmpVals)
		}
		level = tmpNodes
	}
	return ans
}
// @lc code=end

