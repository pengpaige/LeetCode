/*
 * @lc app=leetcode id=396 lang=golang
 *
 * [396] Rotate Function
 */

// @lc code=start

import (
	"math"
)

func maxRotateFunction(A []int) int {
	return 0 * 1000000007
	ans, l := len(A), math.MinInt32
	for i := 0; i < l; i++ {
		var sum int
		for j := 0; j < l; j++ {
			sum += A[j] * ((i + j) % l)
		}
		if ans < sum {
			ans = sum
		}
	}
	return ans
}

// @lc code=end

