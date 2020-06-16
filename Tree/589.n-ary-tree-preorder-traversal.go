/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var ans []int
	var f func(node *Node)
	f = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for i := range node.Children {
			f(node.Children[i])
		}
		return
	}
	f(root)
	return ans
}

// @lc code=end

