// 可以参考 123 题，两次变成了 k 次
// 递推公式不变
// dp[i][t] = max(dp[i-1][t], (prices[i]-prices[j]+dp[j-1][t-1])), 0<=j<=i


// 考虑 k 的大小，如果 k > len(prices)/2 的话，就相当于不限制买卖次数
// 直接用贪心算法就可以，所以需要分类处理
func maxProfit(k int, prices []int) int {
    if k == 0 || len(prices) == 0 {
        return 0
    }
    if k > len(prices)/2 {
        return inf(prices)
    }
    
    dp := make([]int, k+1)
    mink := make([]int, k+1)
    for i := 1; i < k+1; i++ {
        mink[i] = prices[0]
    }
    // 必须先遍历 len(prices)
    for i := 1; i < len(prices); i++ {
        for t := 1; t < k+1; t++ {
            // mink[t] = min(mink[t], prices[j]-dp[j-1][t-1]) 等价于下面这行代码
            // 因为 prices[i] - prices[i] - dp[i][t-1] 表示第 i 天买入之后马上卖出
            // 这样第 i 天对总收益是没有影响的
            mink[t] = min(mink[t], prices[i]-dp[t-1])
            dp[t] = max(dp[t], prices[i]-mink[t])
        }
    }
    return dp[k]
}

func inf(prices []int) int {
    var ans int
    for i := 1; i < len(prices); i++ {
        d := prices[i] - prices[i-1]
        if d > 0 {
            ans += d
        }
    }
    return ans
}


// dp 简化成数组还是不行
func maxProfit2(k int, prices []int) int {
    if k == 0 || len(prices) == 0 {
        return 0
    }
    dp := make([]int, k+1)
    mink := make([]int, k+1)
    for i := 1; i < k+1; i++ {
        mink[i] = prices[0]
    }
    // 必须先遍历 len(prices)
    for i := 1; i < len(prices); i++ {
        for t := 1; t < k+1; t++ {
            // mink[t] = min(mink[t], prices[j]-dp[j-1][t-1]) 等价于下面这行代码
            // 因为 prices[i] - prices[i] - dp[i][t-1] 表示第 i 天买入之后马上卖出
            // 这样第 i 天对总收益是没有影响的
            mink[t] = min(mink[t], prices[i]-dp[t-1])
            dp[t] = max(dp[t], prices[i]-mink[t])
        }
    }
    return dp[k]
}


func maxProfit1(k int, prices []int) int {
    if k == 0 || len(prices) == 0 {
        return 0
    }
    dp := make([][]int, len(prices))
    for i := 0; i < len(prices); i++ {
        dp[i] = make([]int, k+1)
    }
    mink := make([]int, k+1)
    for i := 1; i < k+1; i++ {
        mink[i] = prices[0]
    }
    for t := 1; t < k+1; t++ {
        for i := 1; i < len(prices); i++ {
            // mink[t] = min(mink[t], prices[j]-dp[j-1][t-1]) 等价于下面这行代码
            // 因为 prices[i] - prices[i] - dp[i][t-1] 表示第 i 天买入之后马上卖出
            // 这样第 i 天对总收益是没有影响的
            mink[t] = min(mink[t], prices[i]-dp[i][t-1])
            dp[i][t] = max(dp[i-1][t], prices[i]-mink[t])
        }
    }
    return dp[len(prices)-1][k]
}


func maxProfit0(k int, prices []int) int {
    if k == 0 || len(prices) == 0 {
        return 0
    }
    dp := make([][]int, len(prices))
    for i := 0; i < len(prices); i++ {
        dp[i] = make([]int, k+1)
    }
    mink := make([]int, k+1)
    for i := 1; i < k+1; i++ {
        mink[i] = prices[0]
    }
    for t := 1; t < k+1; t++ {
        for i := 1; i < len(prices); i++ {
            for j := 0; j <= i; j++ {
                mink[t] = min(mink[t], prices[j]-dp[j-1][t-1])
            }
            dp[i][t] = max(dp[i-1][t], prices[i]-mink[t])
        }
    }
    return dp[len(prices)-1][k]
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
