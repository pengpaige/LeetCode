/*
 * @lc app=leetcode id=306 lang=golang
 *
 * [306] Additive Number
 */

// @lc code=start

import (
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	if byte(num[0]) == '0' {
		return false
	}
	for i := 1; i <= len(num)/2; i++ {
		a, _ := strconv.ParseInt(string(num[:i]), 10, 64)
		for j := 1; max(i, j) <= len(num)-i-j; j++ {
			if byte(num[i]) == '0' && j > 1 {
				break
			}
			b, _ := strconv.ParseInt(string(num[i:i+j]), 10, 64)
			return isAdtv(a, b, i+j, num)
		}
	}
	return false
}

func isAdtv(a, b int64, start int, num string) bool {
	if start == len(num) {
		return true
	}
	sum := strconv.FormatInt(a+b, 10)
	idx := strings.Index(num, sum)
	return idx == start && isAdtv(b, a+b, start+len(sum), num)
}

// func helper(num []byte, a, b, lenB, start, end int) bool {
// 	i, _ := strconv.ParseInt(string(num[start:end]), 10, 64)
// 	if i != a+b {
// 		return false
// 	}
// 	// if i == a+b
// 	if end == len(num) {
// 		return true
// 	}
// 	for delta := 1; idx < len(num)-end; idx++ {
// 		if helper(num, b, i, start-end, end, end+delta) {
// 			return true
// 		}
// 	}
// 	return false
// }

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

