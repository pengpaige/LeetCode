/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThreeBruteForce(n int) bool {
	if n <= 0 {
		return false
	}
	for  n % 3 == 0 {
		n /= 3
	}
	if n == 1 {
		return true
	}
    return false
}
// @lc code=end

