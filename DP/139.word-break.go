/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start

// 被同一道题折磨两边并不算什么，当你发现回溯和DP解法答案都参考了同一个人的时候最为致郁
// dp[i] 表示 s 下标 0 ~ i-1 中间的子串符合要求
func wordBreak(s string, wordDict []string) bool {
	sz := len(s)
	if sz == 0 {
		return false
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true // base state
	wm := make(map[string]struct{}, len(wordDict))
	for i := range wordDict {
		wm[wordDict[i]] = struct{}{}
	}

	for i := 1; i <= sz; i++ {
		for j := 0; j < i; j++ {
			_, ok := wm[string(s[j:i])]
			// dp[i] 的含义是从 0 开始到 i-1 的子串满足要求
			// 事实上这样定义 dp[i] 也是最清晰明确方便解题的
			// 下面的 dp[j] = true 会 dp 中记录的子串越来越长最后达到 s 的总长度
			if dp[j] && ok {
				dp[i] = true
				break
			}
		}
	}
	return dp[sz]
}

// @lc code=end

