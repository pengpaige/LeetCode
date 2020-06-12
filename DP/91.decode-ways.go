import "strconv"

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		one, _ := strconv.Atoi(string(s[i]))
		if !inA2Z(one) {
			return false
		}
		dp[i] += dp[i-1]
		two, _ := strconv.Atoi(string(s[i-1 : i+1]))
		if !inA2Z(two) {
			return false
		}
		dp[i] += dp[i-2]
	}
	return dp[len(s)-1]
}

func inA2Z(i int) bool {
	if i >= 1 && i <= 26 {
		return true
	}
	return false
}

// @lc code=end

