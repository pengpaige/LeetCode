func maxProfit1(prices []int) int {
    var ans int
    for i := 0; i < len(prices); i++ {
        for j := i+1; j < len(prices); j++ {
            if tmp := prices[j] - prices[i]; tmp > ans {
                ans = tmp
            }
        }
    }
    return ans
}


// DP 解法
// 第 i 天卖出的最大利润应该是第 i-1 天卖出的最大利润与第 i 天卖出的利润的最大值
// minPrice 代表目前为止发现的最低价
// maxProfit 代表最大利润
func maxProfit(prices []int) int {
    if len(prices) < 1 {
        return 0
    }
    var minPrice, maxProfit int
    minPrice = prices[0]
    for i := 1; i < len(prices); i++ {
        minPrice = min(minPrice, prices[i])
        maxProfit = max(maxProfit, prices[i]-minPrice)
    }
    return maxProfit
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


