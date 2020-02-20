func minimumTotal(triangle [][]int) int {
    h := len(triangle)
    // 初始化二维 slice
    dp := make([][]int, len(triangle))
    for i, _ := range triangle {
        dp[i] = make([]int, len(triangle[i]))
        // 初始化 dp 最后一行
        if i == len(triangle)-1 {
            for j, n := range triangle[i] {
                dp[i][j] = n
            }
        }
    }
    for i := h-2; i>=0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
        }
    }
    return dp[0][0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
