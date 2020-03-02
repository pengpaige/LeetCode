// 需要注意的一点，题目要求必须在卖出之后才能买入，不能同时买入或者卖出两笔
// dp[i][k] = max(dp[i-1][k], (prices[i] - prices[j] + dp[j-1][k-1])); 0 < j < i
// 另外需要注意的是 mink 是表示第 k 次交易的对应值，不是第 i 天的对应值，这个需要自己理解一下
/*
Runtime: 1172 ms, faster than 5.62% of Go online submissions for Best Time to Buy and Sell Stock III.
Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Best Time to Buy and Sell Stock III.
*/
func maxProfit0(prices []int) int {
    if len(prices) <= 0 {
        return 0
    }
    // dp[i][k] 表示第 i 天进行最多 k 次交易时的最大收益
    dp := make([][]int, len(prices))
    for i, _ := range dp {
        dp[i] = make([]int, 3)
    }
    mink := []int{0, prices[0], prices[0]}
    for i := 1; i < len(prices); i++ {
        for k := 1; k < 3; k++ {
            for j := 1; j < i; j++ {
                mink[k] = min(mink[k], prices[j] - dp[j-1][k-1])
            }
            dp[i][k] = max(dp[i-1][k], prices[i] - mink[k])
        }
    }
    return dp[len(prices)-1][2]
}


/*
对上面最内层循环进行优化，因为每次都要从0到i-1循环，其实每次只要比较当前dp[i][k-1]和前面所有一直的dp[i][k-1]即可
Runtime: 4 ms, faster than 95.51% of Go online submissions for Best Time to Buy and Sell Stock III.
Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Best Time to Buy and Sell Stock III.
*/
func maxProfit1(prices []int) int {
    if len(prices) <= 0 {
        return 0
    }
    // dp[i][k] 表示第 i 天进行最多 k 次交易时的最大收益
    dp := make([][]int, len(prices))
    for i, _ := range dp {
        dp[i] = make([]int, 3)
    }
    mink := []int{0, prices[0], prices[0]}
    for i := 1; i < len(prices); i++ {
        for k := 1; k < 3; k++ {
            mink[k] = min(mink[k], prices[i] - dp[i-1][k-1])
            dp[i][k] = max(dp[i-1][k], prices[i] - mink[k])
        }
    }
    return dp[len(prices)-1][2]
}


// 复杂度和上面相同，只是把 dp[i-1][k-1] 换成 dp[i][k-1] 方便统一表达方式
// 另外第 i-1 天如果是最大收益的话，那么 k 不变的情况下，dp[i-1][k-1] == dp[i][k-1]
// 因为第 i 天的收益还是相同 k 次交易的最大收益
func maxProfit2(prices []int) int {
    if len(prices) <= 0 {
        return 0
    }
    // dp[i][k] 表示第 i 天进行最多 k 次交易时的最大收益
    dp := make([][]int, len(prices))
    for i, _ := range dp {
        dp[i] = make([]int, 3)
    }
    mink := []int{0, prices[0], prices[0]}
    for i := 1; i < len(prices); i++ {
        for k := 1; k < 3; k++ {
            mink[k] = min(mink[k], prices[i] - dp[i][k-1])
            dp[i][k] = max(dp[i-1][k], prices[i] - mink[k])
        }
    }
    return dp[len(prices)-1][2]
}






func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
