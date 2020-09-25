/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
    if len(nums) < 2 {
		return 1
	}
	var ans int
	dp := make([]int, len(nums))
	for i, _ : range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := i-1; j >=0; j-- {
			if mums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

