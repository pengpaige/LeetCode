// 花园里有 N 个花盆，每个花盆里都有一朵花。这 N 朵花会在 N 天内依次开放，每天有且仅有一朵花会开放并且会一直盛开下去。

// 给定一个数组 flowers 包含从 1 到 N 的数字，每个数字表示在那一天开放的花所在的花盆编号。

// 例如， flowers[i] = x 表示在第 i+1 天盛开的花在第 x 个花盆中，i 和 x 都在 1 到 N 的范围内。

// 给你一个整数 k，请你输出在哪一天恰好有两朵盛开的花，他们中间间隔了 k 朵花并且都没有开放。

// 如果不存在，输出 -1。

// 样例1

// 输入: 
// flowers: [1,3,2]
// k: 1
// 输出: 2
// 解释: 在第二天，第一朵和第三朵花都盛开了。

// 样例2

// 输入: 
// flowers: [1,2,3]
// k: 1
// 输出: -1

// 注释 :
// 给定的数组范围是 [1, 20000]。

// 双指针
// 实现比较简单，但是理解起来不太容易
// 核心思路就是翻转 bulbs 的映射到 花盆->开花天数 的 days
// 题目要求的「他们中间间隔了 k 朵花并且都没有开放」等价于 days[i] 和 days[i+K+1] 中间的所有有的 days[x] 都比 days[i] 和 days[i+K+1] 大
// 也就是开花时间比两侧 days[i] 和 days[i+K+1] 都晚
// 就能保证 days[i] 和 days[i+K+1] 中间有 K 个花盆没开花
func kEmptySlots(bulbs []int, K int) int {
    if len(bulbs) == 0 {
        return -1
    }
    days := make([]int, len(bulbs))
    for i, n := range bulbs {
        days[n-1] = i
    }

    ans, left, right := 1<<31-1, 0, K+1
    for right < len(days) {
        breakFlag := false
        for i := left+1; i < right; i++ {
            if days[i] < days[left] || days[i] < days[right] {
                left, right = i, i+K+1
                breakFlag = true
                break
            }
        }
        if ! breakFlag {
            ans = min(ans, max(days[left], days[right]))
            left, right = right, right+K+1
        }
    }
    if ans == 1<<31-1 {
        return -1
    }
    return ans+1
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
