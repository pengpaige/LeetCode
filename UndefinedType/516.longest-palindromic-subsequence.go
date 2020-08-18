/*
 * @lc app=leetcode id=516 lang=golang
 *
 * [516] Longest Palindromic Subsequence
 */

// @lc code=start

// 没错我就是抄答案的 https://leetcode-cn.com/problems/longest-palindromic-subsequence/solution/zi-xu-lie-wen-ti-tong-yong-si-lu-zui-chang-hui-wen/
// dp[i][j] 代表 s 的下标冲 i 到 j 范围内的回文子序列长度
// base case: dp[i][i] 代表一个字符一定是回文子序列
func longestPalindromeSubseq(s string) int {
	l := len(s)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		dp[i][i] = 1
	}

	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

