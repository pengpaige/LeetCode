/*
 * @lc app=leetcode id=319 lang=golang
 *
 * [319] Bulb Switcher
 */

// @lc code=start

// 这个规律这辈子也是不可能总结出来的
func bulbSwitch(n int) int {
    return mySqrt(n)
}


func mySqrt(n int) int {
	l, r := 0, n
	for l <= r {
		mid := (l+r)/2
		s := mid*mid
		if mid*mid == n {
			return mid
		} else if s < n {
			if tmp := (mid+1)*(mid+1); tmp > n {
				return mid
			}
			l++
		} else {
			r--
		}
	}
	return l
}
// @lc code=end

