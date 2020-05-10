/*
 * @lc app=leetcode id=390 lang=golang
 *
 * [390] Elimination Game
 */

// @lc code=start


// 对于 math 这个 tag 里的题目
// 我大多数时候是在看完别人的答案之后才彻底搞明白这道题到底想让我干什么
// 这些题目想干的事情无非就是一件 侮辱我的智商
// 如果我有智商的话
func lastRemaining(n int) int {
	first, left, factor, even := 1, true, 1, (n&1 == 0)
	for n > 1 {
		if left || !even {
			first += factor
		}
		left = !left
		factor <<= 1
		n >>= 1
		even = (n&1 == 0)
	}
	return first
}
// @lc code=end

