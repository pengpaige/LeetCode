/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
import (
	"math"
)

func myAtoi(str string) int {
    if len(str) == 0 {
        return 0
    }
	idx, sign, ans := 0, 1, 0
    for idx < len(str) && str[idx] == ' ' {
		idx++
	}
	if idx < len(str) && str[idx] == '-' {
		sign = -1
		idx++
	} else if idx < len(str) && str[idx] == '+' {
        idx++
    }
	for idx < len(str) {
		if str[idx] > '9' || str[idx] < '0' {
			if ans == 0 {
				return 0
			} else {
				return sign*ans
			}
		}
		ans = ans*10 + int(str[idx]-'0')
		if sign > 0 && ans > math.MaxInt32 {
			return math.MaxInt32
        } else if sign < 0 && sign * ans < math.MinInt32 {
			return math.MinInt32
		}
		idx++
	}
	return sign*ans
}
// @lc code=end

