// ans 每一行代表 i+1 个硬币的情况下可能出现的总金额
// 判断最早出现 amount 的一行就是最少的满足要求的硬币个数(需要加一即返回 i+1)
func coinChange1(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    if len(coins) < 1 {
        return -1
    }
    ans := make([][]bool, amount)
    for i, _ := range ans {
        ans[i] = make([]bool, amount+1)
    }
    // 初始化第一行
    var flag bool
    for _, n := range coins {
        if n <= amount {
            flag = true
            ans[0][n] = true
        }
    }
    // 没有一个 coin 比 amount 小，直接报错
    if !flag {
        return -1
    }
    // amount == 1 时
    if ans[0][amount] {
        return 1
    }
    for i := 1; i < amount; i++ {
        for j := 0; j <= amount; j++ {
            if !ans[i-1][j] {
                continue
            }
            for _, c := range coins {
                if j+c <= amount {
                    ans[i][j+c] = true
                }
            }
            if ans[i][amount] {
                return i+1
            }
        }
    }
    return -1
}


// 空间复杂度更小的一种方法
// ans[n] 的值代表找零为 n 时需要的最小硬币数量
// ans 长度设置为 amount+1 是为了索引和找零的数字对齐
// ans 初始化为 amount+1 是因为动态规划的状态数组 ans 中的元素值不可能大于 amount
// 所以对于本题来说 amount+1 相当于无穷大了
func coinChange(coins []int, amount int) int {
    ans := make([]int, amount+1)
    for i := 0; i < amount+1; i++ {
        ans[i] = amount+1
    }
    // 初始化 amount 为 0 的情况，也是 amount 非 0 时的初始条件
    ans[0] = 0
    for i := 0; i <= amount; i++ {
        for _, c := range coins {
            if i < c {
                continue
            }
            ans[i] = min(ans[i], 1+ans[i-c])
        }
    }
    if ans[amount] == amount+1 {
        return -1
    }
    return ans[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
