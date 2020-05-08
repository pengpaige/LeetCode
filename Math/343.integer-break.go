/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start

// 学会数学只能让代码更快，学会递归却能让你找工作更快

var mem = map[int]int{}

func integerBreak(n int) int {
	if m, ok := mem[n]; ok {
		return m
	}
	ans := -1
	for i := 1; i < n; i++ {
		ans = max(ans, max(i*(n-i), i*integerBreak(n-i)))
	}
	mem[n] = ans
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

