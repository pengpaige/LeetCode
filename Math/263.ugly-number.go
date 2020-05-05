/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */

// @lc code=start

// brute force
// 1012/1012 cases passed (4 ms)
// Your runtime beats 35.92 % of golang submissions
// Your memory usage beats 50 % of golang submissions (2.2 MB)
func isUgly(num int) bool {
    if num == 1 {
		return true
	}
	var lastV int
	for {
		lastV = num 
		if num % 2 == 0 {
			num /= 2
		}
		if num % 3 == 0 {
			num /= 3
		}
		if num % 5 == 0 {
			num /= 5
		}
		if num == lastV {
			break
		}
	}
	if num == 1 {
		return true
	}
	return false
}
// @lc code=end

