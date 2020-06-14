/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// @lc code=start

// dp[i][j] 表示 s 和 p 中长度为 i 和 j 的前缀子串能匹配
// 平平无奇动态规划 12 ms, faster than 75.47%
func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	dp := make([][]bool, ls+1)
	for i := range dp {
		dp[i] = make([]bool, lp+1)
	}
	dp[0][0] = true
	for j := 0; j < lp; j++ {
		if p[j] == '*' {
			dp[0][j+1] = true
		} else {
			break
		}
	}

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if p[j-1] == '*' {
				// * 匹配空串，* 匹配任意一个字符，* 匹配多个字符
				dp[i][j] = dp[i][j-1] || dp[i-1][j-1] || dp[i-1][j]
			}
			if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[ls][lp]
}

// 回溯+记忆 316 ms, faster than 16.04%
func isMatch_Backtracking(s string, p string) bool {
	ls, lp := len(s), len(p)
	mem := make(map[[2]int]bool, 0) //不加记忆就超时
	var f func(sc, pc int) bool
	f = func(sc, pc int) (ans bool) {
		if ret, ok := mem[[2]int{sc, pc}]; ok {
			return ret
		}

		// 我竟然写出了 Golang 风格的记忆化！
		// 这里注意一下 ans 和 mem 都是通过闭包方式引用传递
		defer func() {
			mem[[2]int{sc, pc}] = ans
		}()

		if sc == ls && pc == lp {
			ans = true
			return
		}
		if pc == lp {
			ans = false
			return
		}
		if p[pc] == '*' {
			// * 匹配空串，* 匹配任意一个字符，* 匹配多个字符
			if sc < ls {
				ans = f(sc, pc+1) || f(sc+1, pc+1) || f(sc+1, pc)
			} else {
				ans = f(sc, pc+1)
			}
			return
		}
		if sc < ls && (p[pc] == '?' || p[pc] == s[sc]) {
			ans = f(sc+1, pc+1)
			return
		}
		return
	}
	return f(0, 0)
}

// @lc code=end

