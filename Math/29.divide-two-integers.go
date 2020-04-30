/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
import (
	"math"
)


func divide(dividend int, divisor int) int {
	var positive bool
	if (dividend >= 0 && divisor > 0) || (dividend <= 0 && divisor < 0) {
		positive = true
	} else {
		positive = false
	}
	absDividend, absDivisor, cmp, ans, currentQ := abs(dividend), abs(divisor), abs(divisor), 0, 1
	if absDividend < absDivisor {
		return 0
	}
	for (cmp << 1) <= absDividend {
		cmp <<= 1
		currentQ <<= 1
	}
	for absDividend >= absDivisor {
		absDividend -= cmp
		ans += currentQ
		for cmp > absDividend {
			cmp >>= 1
			currentQ >>= 1
		}
	}
	if positive {
		if ans > math.MaxInt32 {
			return math.MaxInt32
		}
		return ans
	} else {
		ans *= -1
		if ans < math.MinInt32 {
			return math.MaxInt32
		}
		return ans
	}
}


func abs(i int) int {
	if i < 0 {
		return -1*i
	}
	return i
}
// @lc code=end

