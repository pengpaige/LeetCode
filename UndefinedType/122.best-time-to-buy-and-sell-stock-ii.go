/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start

// DP 递推公式看下面代码
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// 左下标代表第 i 天
	// 右下标 0 代表本次操作之后持有现金, 右下标 1 代表本次操作后持有股票
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = 0, -1*prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	// 最后一定是股票卖出之后持有现金最多
	return dp[len(prices)-1][0]
}

// BruteForce DFS Time Limit Exceeded(198/200 cases passed)
func maxProfit_DFS(prices []int) int {
	maxProfit := 0
	var dfs func(idx, profit, status int)
	dfs = func(idx, profit, status int) {
		if idx == len(prices) {
			maxProfit = max(maxProfit, profit)
			return
		}

		// 不操作
		dfs(idx+1, profit, status)

		// 根据 status 判断买入或者卖出, PS 题目要求买入之后必须卖出才能再买入
		if status == 0 { // 本次操作之前持有现金, 可以买入
			dfs(idx+1, profit-prices[idx], 1)
		} else if status == 1 { // 本次操作之前持有股票, 可以卖出
			dfs(idx+1, profit+prices[idx], 0)
		}
	}

	dfs(0, 0, 0)
	return maxProfit
}

// 贪心算法 时间复杂度 O(n)
func maxProfit_Greedy(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var ans int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			ans += prices[i+1] - prices[i]
		}
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

