/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 首先这是一棵 BST 祖先节点的值
// 利用 BST 的性质，祖先节点的 Val 介于 p.Val q.Val 之间
// 反之，如果不介于二者之间的值是 p q 的祖先节点， 那么一定存在更底层的祖先节点 
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	node := root
	for !in(node, p, q) {
		if node == nil {
			break
		}
		if node.Val > max(p, q).Val {
			node = node.Left
		}
		if node.Val < min(p, q).Val {
			node = node.Right
		}
	}
	if node != nil {
		ans = node
	}
	return ans
}

func in(node, p, q *TreeNode) bool {
	return node.Val >= min(p, q).Val && node.Val <= max(p, q).Val 
}

func min(p, q *TreeNode) *TreeNode {
	if p.Val < q.Val {
		return p
	}
	return q
}

func max(p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		return p
	}
	return q
}
// @lc code=end

