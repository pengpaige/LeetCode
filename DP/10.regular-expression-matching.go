/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start

// 测试例中不允许 * 号开头，可以不考虑这种 corner case
// 而且实际上以 * 号开头也不符合 * 号本身的定义

// Backtracking
func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	var f func(sc, pc int) bool
	f = func(sc, pc int) bool {
		if sc == ls {
			if pc == lp || (pc == lp-1 && p[pc] == '*') {
				return true
			}
			return false
		}
		if p[pc] == '.' {
			return f(sc+1, pc+1)
		} else if p[pc] == '*' {
			// 正常情况 * 号不会出现在第一个字符位置
			if pc > 0 && p[pc-1] == s[sc] {
				return f(sc+1, pc)
			}
			return f(sc, pc+1)
		} else if p[pc] == s[sc] {
			return f(sc+1, pc+1)
		}
		return false
	}
	return f(0, 0)
}

// @lc code=end

