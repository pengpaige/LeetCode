/*
 * @lc app=leetcode id=429 lang=golang
 *
 * [429] N-ary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	var ans [][]int
	if root == nil {
		return nil
	}
	var level []*Node
	level = append(level, root)
	for len(level) > 0 {
		var l []int
		var tmp []*Node
		for _, node := range level {
			l = append(l, node.Val)
			if node.Children != nil {
				tmp = append(tmp, node.Children...)
			}
		}
		ans = append(ans, l)
		level = tmp
	}
	return ans
}
// @lc code=end

