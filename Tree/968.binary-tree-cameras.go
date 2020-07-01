/*
 * @lc app=leetcode id=968 lang=golang
 *
 * [968] Binary Tree Cameras
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

// 逻辑很复杂但是实现起来很简单
// 贪心的字底向上遍历二叉树，同时只在必须放 camera 的时候才放
func minCameraCover(root *TreeNode) int {
	var ans int
	covered := make(map[*TreeNode]bool, 0)
	// 避免因为子节点是 nil 导致误以为该子节点没有被标记
	covered[nil] = true
	var postOrder func(curr, par *TreeNode)
	postOrder = func(curr, par *TreeNode) {
		if curr == nil {
			return
		}
		postOrder(curr.Left, curr)
		postOrder(curr.Right, curr)
		// 实际上递归只递归到倒数第二层的节点
		// 但是为了代码的简洁没有在递归到下一层的时候直接限制
		// 因为直接限制的话需要取子节点的子节点，太繁琐
		// 这里将这个逻辑转换为只对父节点为 nil 的节点即 root 判断当前值是否被标记
		// 因为字底向上遍历的过程中，下一层节点被放置了 camera 之后本层节点也会被覆盖不需要再放 camera
		// 也就是只有当前节点 curr 的子节点未被标记时才需要在 curr 位置放 camera
		// 但由于 root 没有上层节点，即使 root 子节点被标记但自身未被覆盖时，也需要单独标记 root
		if (par == nil && !in(curr, covered)) || !in(curr.Left, covered) || !in(curr.Right, covered) {
			ans++
			covered[par] = true
			covered[curr] = true
			covered[curr.Left] = true
			covered[curr.Right] = true
		}
	}
	postOrder(root, nil)
	return ans
}

func in(node *TreeNode, m map[*TreeNode]bool) bool {
	_, ok := m[node]
	return ok
}
// @lc code=end

