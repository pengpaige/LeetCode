/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	mid, s, l, r := 0, 0, 0, x
	// (x/2)^2 = x 在离散的情况下是 x = 5 时成立
	// 当 x > 5 时，(x/2)^2 > x
	// 所以使用下面的方式来减少循环次数
	if x > 5 {
		r = x / 2
	}
	for l <= r {
		mid = (l + r) / 2
		s = mid * mid
		if s == x {
			return mid
		} else if s > x {
			r = mid - 1
		} else if s < x {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			l = mid + 1
		}
	}
	return -1
}

// @lc code=end

