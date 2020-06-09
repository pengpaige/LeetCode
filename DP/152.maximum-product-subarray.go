/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	ans, curr := -1<<31, 1
	for _, n := range nums {
		if n == 0 {
			curr = 1
			continue
		}
		curr *= n

	}
}
// @lc code=end

