/*
 * @lc app=leetcode id=397 lang=golang
 *
 * [397] Integer Replacement
 */

// @lc code=start


var mem = make(map[int]int, 0)
func integerReplacement(n int) int {
    return dfs(n)
}


func dfs(n int) int {
	if count, ok := mem[n]; ok {
		return count
	}
	if n <= 3 {
		return n-1
	}
	if n%2 == 0 {
		return dfs(n/2) + 1
	}
	ret := min(dfs(n+1), dfs(n-1)) + 1
	mem[n] = ret
	return ret
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

