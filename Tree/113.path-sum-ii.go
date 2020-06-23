/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, sum int) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var path []int
	var findSum func(node *TreeNode, curr int)
	findSum = func(node *TreeNode, curr int) {
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil {
			if curr+node.Val == sum {
				tmp := make([]int, len(path))
				copy(tmp, path)
				ans = append(ans, tmp)
				return
			}
			return
		}
		if node.Left != nil {
			findSum(node.Left, curr+node.Val)
		}
		if node.Right != nil {
			findSum(node.Right, curr+node.Val)
		}
	}
	findSum(root, 0)
	return ans
}

// @lc code=end

