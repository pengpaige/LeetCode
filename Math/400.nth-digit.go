/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 */

// @lc code=start

// Time Limit Exceeded
// 34/70 cases passed (N/A)
// Testcase
// 1000000
// import "strconv"
// func findNthDigit(n int) int {
// 	s := ""
// 	for i := 0; i <= n; i++ {
// 		s += strconv.Itoa(i)
// 	}
// 	i, _ := strconv.Atoi(string(s[n]))
// 	return i
// }

// https://dwz1.cc/Txvm3nf
import (
	"fmt"
	"math"
)

func findNthDigit(n int) int {
	n--
	for digitCount := 1; digitCount < 11; digitCount++ {
		firstNumber := int(math.Pow10(digitCount - 1))
		currMax := 9 * firstNumber * digitCount
		if n < currMax {
			return int(fmt.Sprintf("%d", firstNumber+n/digitCount)[n%digitCount] - '0')
		}
		n -= currMax
	}
	return 0
}

// @lc code=end

