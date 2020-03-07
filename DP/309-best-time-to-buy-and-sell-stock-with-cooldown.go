/*
0 不持股 非冷冻期，即卖出日期的隔天
1 持股 今天刚买入
2 冷冻期就是指卖出日期后一天
不持股，前一天不持股，或者前一天是冷冻期            dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][2]); 
持股，前一天持股，或者前一天是不持股但今天买入      dp[i][1] = Math.max(dp[i - 1][1], dp[i - 1][0] - prices[i]); 
冷冻期，只能是前一天是持股                      dp[i][2] = dp[i - 1][1] + prices[i];
*/
func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, 3)
    }
    dp[0][0] = 0
    dp[0][1] = -prices[0]
    dp[0][2] = 0
    for i := 1; i < n; i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][2])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
        dp[i][2] = dp[i-1][1] + prices[i]
    }
    return max(dp[n-1][0], dp[n-1][2])
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
