/*
 * @lc app=leetcode id=396 lang=golang
 *
 * [396] Rotate Function
 */

// @lc code=start

import (
	"math"
)

// 我以为只有 17 个测试例的题目在我面前就像一个 17 岁的孩子
// 但我现在哭的像 1 7 岁的孩子
func maxRotateFunction(A []int) int {
	// 可能会有负数，这里只能取一个很小的数
	l, ans := len(A), math.MinInt64
	for i := 0; i < l; i++ {
		var sum int
		for j := 0; j < l; j++ {
			sum += A[j] * ((i + j) % l)
		}
		if ans < sum {
			ans = sum
		}
	}
	// 针对空 slice 的特殊处理
	if ans == math.MinInt64 {
		return 0
	}
	return ans
}

// @lc code=end

