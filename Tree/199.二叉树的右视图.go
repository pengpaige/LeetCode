/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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

// 层序遍历然后每层取队列最后一个值放到结果中
func rightSideView_LevelTraverse(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		// 每层把最右边的值放入结果中
		ans = append(ans, q[len(q)-1].Val)
		newq := make([]*TreeNode, 0)
		for len(q) != 0 {
			curr := q[0]
			if curr.Left != nil {
				newq = append(newq, curr.Left)
			}
			if curr.Right != nil {
				newq = append(newq, curr.Right)
			}
			q = q[1:]
		}
		q = newq
	}
	return ans
}

// DFS
// 难点在于需要记录当前访问的层数
// 按照 根-右-左 的顺序遍历二叉树时, 每层第一个访问到的节点就是结果集中的值
// ps: 深度 depth 从 0 开始
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		// depth 从 0 开始, ans 的长度也是从 0 开始, 刚好匹配上
		if len(ans) == depth {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return ans
}

// @lc code=end

