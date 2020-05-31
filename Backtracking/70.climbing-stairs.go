/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

// Your runtime beats 100 % of golang submissions
// Your memory usage beats 22.22 % of golang submissions (2 MB)
func climbStairs_Recursion(n int) int {
	var rec func(n int) int
	mem := make(map[int]int)
	rec = func(n int) int {
		// 不加记忆就超时
		if _, ok := mem[n]; ok {
			return mem[n]
		}
		if n < 2 {
			return 1
		}
		ret := rec(n-1) + rec(n-2)
		mem[n] = ret
		return ret
	}
	return rec(n)
}

// Your runtime beats 100 % of golang submissions
// Your memory usage beats 22.22 % of golang submissions (2 MB)
func climbStairs_DP(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n]
}

// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)
func climbStairs_DP_OPT(n int) int {
	if n < 2 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
// @lc code=end