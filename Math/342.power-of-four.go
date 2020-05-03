/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
func isPowerOfFour(num int) bool {
	return num > 0 &&
		num & (num - 1) == 0 &&
		num & 0xaaaaaaaa == 0
}
// @lc code=end

