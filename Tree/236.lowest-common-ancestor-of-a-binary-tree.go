/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// dfs 函数里的逻辑没办法包含 p q 互为父子的情况
	// 这里只能单独判断
    if isAcst(p, q) {
        return p
    }
    if isAcst(q, p) {
        return q
    }
    var ans *TreeNode
    var dfs func(node, p, q *TreeNode) (bool, bool)
	dfs = func(node, p, q *TreeNode) (bool, bool) {
		if node == nil {
			return false, false
		}

		leftFoundP, leftFoundQ := dfs(node.Left, p, q)
		rightFoundP, rightFoundQ := dfs(node.Right, p, q)
		foundP, foundQ := leftFoundP || rightFoundP, leftFoundQ || rightFoundQ
		// 因为是后序遍历，暂时想不到很好的剪枝办法
		if ans == nil && (foundP && foundQ) {
			ans = node
			return true, true
		}

		if node == p {
			foundP = true
		}
		if node == q {
			foundQ = true
		}
		return foundP, foundQ
	}
	fp, fq := dfs(root, p, q)
    if ans == nil && (fp && fq) {
		ans = root
	}
	return ans
}

// return p if p is ancestor of q
func isAcst(p, q *TreeNode) bool {
    var found bool
    var find func(node *TreeNode)
    find = func(node *TreeNode) {
        if node == nil {
            return
        }
        if found {
            return
        }
        if node == q {
            found = true
            return
        }
        find(node.Left)
        find(node.Right)
    }
    find(p)
    return found
}


// @lc code=end

