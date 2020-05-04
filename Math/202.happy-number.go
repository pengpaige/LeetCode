/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start

// 感觉这对我来说并不是一道简单题
// 这次还是想不起来细节的话就看这个吧 https://dwz1.cc/jHApVSr
func isHappy(n int) bool {
	m := make(map[int]bool, 0)
	// 当然我也是不可能自己写出这种代码的
	for ; n != 1 && !m[n]; n, m[n] = digitPowSum(n), true {
		// 这辈子都不可能
	}
	return n == 1
}

func digitPowSum(n int) int {
	var sum int
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// @lc code=end

