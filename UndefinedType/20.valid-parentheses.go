/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	var stk []byte
	for i := 0; i < len(s); i++ {
		if len(stk) == 0 || !match(stk[len(stk)-1], s[i]) {
			stk = append(stk, s[i])
        } else if match(stk[len(stk)-1], s[i]) {
			stk = stk[:len(stk)-1]
		}
	}
	if len(stk) != 0 {
		return false
	}
	return true
}

// left 和 right 顺序必须和下面的 if else 一致
// 否则不能配对
func match(left, right byte) bool {
    if left == '(' && right == ')' {
        return true
    }
    if left == '[' && right == ']' {
        return true
    }
    if left == '{' && right == '}' {
        return true
    }
    return false
}
// @lc code=end

