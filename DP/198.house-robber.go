/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	sz := len(nums)
	switch sz {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	// dp[i] 代表截止到 nums[i] 满足要求的最大子序列和
	dp := make([]int, sz)
	dp[0] = nums[0]
	dp[1] = nums[1]
	ans := max(dp[0], dp[1])
	for i := 2; i < sz; i++ {
		// 有可能间隔 1 个数字和不是最大比如 [2,1,1,2] 所以需要遍历找最大
		maxbf := 0
		for j := i - 2; j >= 0; j-- {
			maxbf = max(maxbf, dp[j])
		}
		dp[i] = maxbf + nums[i]
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

