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
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		one, _ := strconv.Atoi(string(s[i-1]))
		if inA2Z(one) && one != 0 {
			dp[i] += dp[i-1]
		}

		if i < 2 {
			continue
		}

		if s[i-1-1] == '0' {
			if s[i-1] == '0' {
				return 0
			}
			continue
		}

		two, _ := strconv.Atoi(string(s[i-1-1 : i]))
		if inA2Z(two) {
			dp[i] += dp[i-2]
		}

	}
	return dp[len(s)]
}

func inA2Z(i int) bool {
	if i >= 1 && i <= 26 {
		return true
	}
	return false
}

// @lc code=end

