/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	wd := make(map[string]struct{}, len(wordDict))
	for _, s := range wordDict {
		wd[s] = struct{}{}
	}
	return breakable(s, wd)
}

func breakable(s string, wd map[string]struct{}) bool {
	mem := make(map[int]bool, len(wd))
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if _, ok := mem[start]; ok {
			return mem[start]
		}
		if start == len(s) {
			return true
		}
		for i := start; i < len(s); i++ {
			_, ok := wd[s[start:i+1]]
			if ok && dfs(i+1) {
				mem[start] = true
				return true
			}
		}
		mem[start] = false
		return false
	}
	return dfs(0)
}

// @lc code=end

