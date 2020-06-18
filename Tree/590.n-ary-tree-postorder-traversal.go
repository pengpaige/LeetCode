/*
 * @lc app=leetcode id=590 lang=golang
 *
 * [590] N-ary Tree Postorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// Iteration
// 如果忘了怎么写递归后序遍历可以看 149 题
func postorder(root *Node) []int {
	var ans []int
	var stack []*Node
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			stack = append(stack, top, nil)
			for i := len(top.Children) - 1; i >= 0; i-- {
				stack = append(stack, top.Children[i])
			}
		} else if len(stack) > 0 {
			top := stack[len(stack)-1]
			ans = append(ans, top.Val)
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}

// @lc code=end

