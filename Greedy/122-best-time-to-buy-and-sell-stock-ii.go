// 贪心算法 时间复杂度 O(n)
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    var ans int
    for i := 0; i < len(prices)-1; i++ {
        if prices[i] < prices[i+1] {
            ans += prices[i+1]-prices[i]
        }
    }
    return ans
}
