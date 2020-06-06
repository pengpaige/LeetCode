/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
// 这不是二维 DP 这是二维迷宫
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] 表示 text1 前 i 个字符 和 text2 中前 j 个字符的 LCS 长度
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 0; i <= len(text1); i++ {
		for j := 0; j <= len(text2); j++ {
			if i == 0 || j == 0{
				dp[i][j] = 0
				continue
			}
			// 定义 dp 的时候图方便直接用下标 i 和 j 表示原始字符串的长度
			// 下面求原始串字符的时候就得减 1
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1]+1
			} else{
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
// @lc code=end

