/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, ans := balanced(root)
    return ans
}

func balanced(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    l, lok := balanced(root.Left)
    r, rok := balanced(root.Right)
    if lok && rok && abs(l-r) <=1 {
        if l > r {
            return l+1, true
        }
        return r+1, true
    }
    return 0, false
}

func abs(x int) int {
    if x >= 0 {
        return x
    }
    return -1*x
}
