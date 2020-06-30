/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
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
func findFrequentTreeSum(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	sums := make([]int, 0)
	var postOrder func(node *TreeNode) int
	postOrder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := postOrder(node.Left)
		rightSum := postOrder(node.Right)
		currSum := leftSum + rightSum + node.Val
		if currSum > max {
			max = currSum
		}
		if currSum < min {
			min = currSum
		}
		sums = append(sums, currSum)
		return currSum
	}
	postOrder(root)
	count := make(map[int]int, 0)
	max := 0
	for _, sum := range sums {
		count[sum]++
		if count[sum] > max {
			max = count[sum]
		}
	}
	for sum, c := range count {
		if c == max {
			ans = append(ans, sum)
		}
	}
	return ans
}

// @lc code=end

