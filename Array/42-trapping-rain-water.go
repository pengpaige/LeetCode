// import "container/heap" 可以成功引用到，但前提是你得记住功能和参数

// brute force
func trap_bf(height []int) int {
    ans := 0
    for i := 0; i < len(height); i++ {
        lm, rm := 0, 0
        for l := i; l >= 0; l-- {
            lm = max(lm, height[l])
        }
        for r := i; r < len(height); r++ {
            rm = max(rm, height[r])
        }
        ans += min(lm, rm) - height[i]
    }
    return ans
}


// 类似 DP 的思想
func trap_dp(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }
    ans := 0
    lm, rm := make([]int, n), make([]int, n)
    lm[0] = height[0]
    for i := 1; i < n; i++ {
        lm[i] = max(lm[i-1], height[i])
    }
    rm[n-1] = height[n-1]
    for i := n-2; i >= 0; i-- {
        rm[i] = max(rm[i+1], height[i])
    }
    for i := 0; i < n; i++ {
        ans += min(rm[i], lm[i]) - height[i]
    }
    return ans
}


/*
two points
不够直观很难想到
除非从上面 DP 的结论基础上进行优化
http://dwz1.cc/kMWzQ3r
*/
func trap(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }
    var ans int
    l, r, lm, rm := 0, n-1, 0, 0
    for l < r {
        if height[l] < height[r] {
            lm = max(lm, height[l])
            ans += lm - height[l]
            l++
        } else {
            rm = max(rm, height[r])
            ans += rm - height[r]
            r--
        }
    }
    return ans
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
