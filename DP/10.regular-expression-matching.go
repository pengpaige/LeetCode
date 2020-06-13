/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start

// 测试例中不允许 * 号开头，可以不考虑这种 corner case
// 而且实际上以 * 号开头也不符合 * 号本身的定义

// DP
func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	// dp[i][j] 表示 s[:i] 能被 p[:j] 匹配，表示的是长度为 i j
	// 也可以理解成 s 和 p 的下标从 1 开始算，0 表示空串
	// 这样设定的原因是什么我也不知道
	dp := make([][]bool, ls+1)
	for i := 0; i <= ls; i++ {
		dp[i] = make([]bool, lp+1)
	}
	// base case
	dp[0][0] = true
	for j := 0; j < lp; j++ {
		if p[j] == '*' && dp[0][j-1] {
			dp[0][j+1] = true
		}
	}
	// j = 0 且 i != 0 肯定是 false
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				if p[j-2] != s[i-1] {
					dp[i][j] = dp[i][j-2]
				}
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j]
				}
			}
			// 既不相等又不是 * 号的话那就是 false 不用设定这样的 dp[i][j]
		}
	}
	return dp[ls][lp]
}

// Backtracking
func isMatch_BackTracking(s string, p string) bool {
	ls, lp := len(s), len(p)
	var f func(sc, pc int) bool
	f = func(sc, pc int) bool {
		if sc == ls && pc == lp {
			return true
		}
		if pc == lp {
			return false
		}
		if p[pc] == '*' {
			var use0P, use1OrMoreP bool
			// 正常情况 * 号不会出现在第一个字符位置
			// sc < ls 这个条件是因为可能 s 已经走到结尾但是 p 还没有到结尾
			// 这时候还是需要继续移动 p 的位置 pc，但不需要再比较 s[sc] 了
			// 因为 sc 结束不是真的结束，需要等到 p 也结束
			if pc > 0 && sc < ls && (p[pc-1] == s[sc] || p[pc-1] == '.') {
				use1OrMoreP = f(sc+1, pc) || f(sc+1, pc+1)
			}
			use0P = f(sc, pc+1)
			return use0P || use1OrMoreP
		}
		var eq, bfStar bool
		if sc < ls && (p[pc] == s[sc] || p[pc] == '.') {
			eq = f(sc+1, pc+1)
		}
		if pc+1 < lp && p[pc+1] == '*' {
			bfStar = f(sc, pc+2)
		}
		return eq || bfStar
	}
	return f(0, 0)
}

// Backtracking with memo 100%
func isMatch_BackTrackingWithMemory(s string, p string) bool {
	ls, lp := len(s), len(p)
	var f func(sc, pc int) bool
	mem := make(map[[2]int]bool, 0)
	f = func(sc, pc int) bool {
		if ret, ok := mem[[2]int{sc, pc}]; ok {
			return ret
		}
		if sc == ls && pc == lp {
			return true
		}
		if pc == lp {
			return false
		}
		if p[pc] == '*' {
			var use0P, use1OrMoreP bool
			// 正常情况 * 号不会出现在第一个字符位置
			if pc > 0 && sc < ls && (p[pc-1] == s[sc] || p[pc-1] == '.') {
				use1OrMoreP = f(sc+1, pc) || f(sc+1, pc+1)
			}
			use0P = f(sc, pc+1)
			mem[[2]int{sc, pc}] = use0P || use1OrMoreP
			return use0P || use1OrMoreP
		}
		var eq, bfStar bool
		if sc < ls && (p[pc] == s[sc] || p[pc] == '.') {
			eq = f(sc+1, pc+1)
		}
		if pc+1 < lp && p[pc+1] == '*' {
			bfStar = f(sc, pc+2)
		}
		mem[[2]int{sc, pc}] = eq || bfStar
		return eq || bfStar
	}
	return f(0, 0)
}

// @lc code=end

