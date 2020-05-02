/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := (l+r)/2
		m2 := mid*mid
		if m2 == num {
			return true
		} else if m2 < num {
			l = mid+1
		} else if m2 > num {
			r = mid-1
		}
	}
	return false
}
// @lc code=end

