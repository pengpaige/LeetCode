/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
// 找不到优化方法了，这辈子都不可能找到了，单纯 DP 过不了测试例 31 32
func wordBreak(s string, wordDict []string) []string {
	sz := len(s)
	if sz == 0 {
		return nil
	}
	wd := make(map[string]struct{}, len(wordDict))
	short, long := 1, 1
	for _, w := range wordDict {
		wd[w] = struct{}{}
		if len(w) > long {
			long = len(w)
		}
		if len(w) < short {
			short = len(w)
		}
	}
	// dp[i] 表示 s 中下标 0~i-1 范围内的字符能拆分成句子(可能有多个所以是 slice)
	dp := make([][]string, sz+1)
	dp[0] = []string{""}
	for i := short; i <= sz; i++ {
		for j := 0; j <= i-short; j++ {
			word := string(s[j:i])
			_, ok := wd[word]
			if len(dp[j]) > 0 && ok {
				for k := 0; k < len(dp[j]); k++ {
					var gap string
					if dp[j][k] != "" {
						gap = " "
					}
					dp[i] = append(dp[i], dp[j][k]+gap+word)
				}
			}
		}
	}
	return dp[sz]
}

// @lc code=end

