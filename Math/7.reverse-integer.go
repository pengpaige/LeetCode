/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
import "math"

func reverse(x int) int {
	var ans int
    for x != 0 {
		ans = 10*ans + x%10
		if ans > math.MaxInt32 || ans < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return ans
}
// @lc code=end

