import "sort"

/*
 * @lc app=leetcode id=987 lang=golang
 *
 * [987] Vertical Order Traversal of a Binary Tree
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

// 别想一步到位，先存储节点位置再按要求输出
func verticalTraversal(root *TreeNode) [][]int {
	var ans [][]int
	mem := make(map[int][][2]int, 0)
	var f func(node *TreeNode, x, y int)
	f = func(node *TreeNode, x, y int) {
		if node == nil {
			return
		}
		mem[x] = append(mem[x], [2]int{y, node.Val})
		if node.Left != nil {
			f(node.Left, x-1, y-1)
		}
		if node.Right != nil {
			f(node.Right, x+1, y-1)
		}
	}
	f(root, 0, 0)
	var keys []int
	for x := range mem {
		m := mem[x]
		// 对 y 排序, 如果 y 相等就按 node.Val 升序
		sort.Slice(m, func(i, j int) bool {
			if m[i][0] == m[j][0] {
				return m[i][1] < m[j][1]
			}
			return m[i][0] > m[j][0]
		})
		keys = append(keys, x)
	}
	// 对 x 排序, 因为 map 的 key 无法直接排序, 需要 keys 做辅助 slice
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		var ret []int
		for _, xy := range mem[k] {
			ret = append(ret, xy[1])
		}
		ans = append(ans, ret)
	}
	return ans
}

// @lc code=end

