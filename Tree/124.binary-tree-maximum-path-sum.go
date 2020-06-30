/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
func maxPathSum(root *TreeNode) int {
	ans := -1 << 31
	var postOrder func(node *TreeNode) int
	postOrder = func(node *TreeNode) int {
		var ret int
		if node == nil {
			return 0
		}
		left := postOrder(node.Left)
		right := postOrder(node.Right)
		bigger := max(left, right)
		// 最大路径和可能「拐弯」也可能是「直」的
		// left 和 right 表示左右分支最大和（这里最大指「直的」最大）
		// 找包含当前节点值得最大和时需要同时考虑「弯的」和「直的」
		// 但是返回给上一层递归的时候只返回「直的」最大和
		if left > 0 && right > 0 {
			ans = max(ans, left+right+node.Val)
		}
		if bigger < 0 {
			ret = node.Val
		} else {
			ret = node.Val + bigger
		}
		ans = max(ans, ret)
		return ret
	}
	postOrder(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

