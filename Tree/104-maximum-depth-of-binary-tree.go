/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    ll := maxDepth(root.Left)
    rl := maxDepth(root.Right)
    if ll >= rl {
        return ll+1
    }
    return rl+1
}
