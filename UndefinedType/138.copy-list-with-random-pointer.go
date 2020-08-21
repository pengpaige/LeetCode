/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node, 0)
	var bt func(node *Node) *Node
	bt = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, ok := m[node]; ok {
			return m[node]
		}
		newNode := &Node{Val: node.Val}
		// 因为实在原表上回溯, 所以需要保存访问过的原表节点而不是新表节点
		m[node] = newNode
		newNode.Next = bt(node.Next)
		newNode.Random = bt(node.Random)
		return newNode
	}
	return bt(head)
}

// @lc code=end

