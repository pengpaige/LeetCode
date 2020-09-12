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

// 20200912
// 之前的解法在判断是否互为父子, 以及具体回溯的实现过程中, 不够简洁
// 下面优化之后的解法, 回溯时递归的子流程不再返回 p q 到底找到哪个
// 只返回本次找到的节点(p 或 q)或者 nil
// 详细讲解可以看[这里](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }
    l := lowestCommonAncestor(root.Left, p, q)
    r := lowestCommonAncestor(root.Right, p, q)
    // 首先, 如果 l 和 r 都非空的话, 一定是 p 和 q 分别在 root 左右子树中
    // 因为, 因为一个节点不可能同时存在于左右子树中
    // 其次, 在第一次找到这样左右子树分别包含 pq 的 root 之后
    // 后续的回溯过程中, 上述第一次找到的 root 会一直被向上 return
    // 直到回溯结束, 其他回溯分支不会再 return 非空的值
    // 所以最终的返回值就是第一次找到的这个左右子树分别包含 pq 的 root
    if l != nil && r != nil {
        return root
    } else if l == nil {
        return r
    } else if r == nil {
        return l
    } else {
        return nil
    }
    // 上面的返回结果逻辑可以再简化为下面的
    // if l == nil {
    //     return r
    // } else if r == nil {
    //     return l
    // } else {
    //     return root
    // }
}


func lowestCommonAncestor_(root, p, q *TreeNode) *TreeNode {
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

        // 因为是找最低公共祖先，所以使用后续遍历先从左右子树里找是最自然的思路
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

