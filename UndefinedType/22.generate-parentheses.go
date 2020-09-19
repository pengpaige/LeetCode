/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	if n == 0 {
		return ans
	}
	// l, r 代表已经使用的左右括号数量
	var dfs func(l, r int, curr string)
	dfs = func(l, r int, curr string) {
		if l == n && r == n {
			ans = append(ans, curr)
			return
		}
		// 在左括号用完之前可能会同时对 curr 增加左括号和右括号
		// 就是下面的两个判断逻辑都为 true
		if l < n {
			dfs(l+1, r, curr+"(")
		}
		if r < l {
			dfs(l, r+1, curr+")")
		}
	}
	dfs(0, 0, "")
	return ans
}

// @lc code=end

