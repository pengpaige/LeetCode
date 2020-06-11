/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	// dp[i][j] == true 代表 s 中 [i, j] 的子串是回文的
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	// 可以不初始化对角线，因为这些不会被后续子问题参考，但一个字符确实是回文
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	start, maxLen := 0, 1
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return string(s[start : start+maxLen])
}

// @lc code=end

