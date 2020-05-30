/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
func wordBreakBy139(s string, wordDict []string) []string {
	var ans []string
	wd := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wd[word] = struct{}{}
	}
	// 因为下面这个测试例的存在，最简单的方式就是先判断 s 能不能按词典拆分成句子
	// 最吊诡的是这个解法竟然能拿双百
	// s :="a...73a...aba...73a...a"
	// wordDict := []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
	if !breakable(s, wd) {
		return ans
	}
	bt140(0, s, "", wd, &ans)
	return ans
}

// 回溯 从前向后
func bt140(start int, s, currStr string, wd map[string]struct{}, ans *[]string) {
	if start == len(s) {
		*ans = append(*ans, currStr)
		return
	}
	newStr := currStr
	for i := start; i < len(s); i++ {
		cs := s[start : i+1]
		if _, ok := wd[cs]; !ok {
			continue
		}
		if start != 0 {
			cs = " " + cs
		}
		newStr += cs
		bt140(i+1, s, newStr, wd, ans)
		newStr = currStr
	}
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

// 回溯 从后向前
func wordBreak(s string, wordDict []string) []string {
	sz := len(wordDict)
	wd := make(map[string]struct{}, sz)
	mem := make(map[int][]string, sz)
	for _, word := range wordDict {
		wd[word] = struct{}{}
	}

	var dfs func(start int) []string
	dfs = func(start int) []string {
		if _, ok := mem[start]; ok {
			return mem[start]
		}
		var ret []string
		if start == len(s) {
			return ret
		}
		for i := start; i < len(s); i++ {
			w := s[start : i+1]
			if _, ok := wd[w]; !ok {
				continue
			}
			sfxs := dfs(i + 1)
			// s 中的最后一个可匹配的 word 就不需要再加空格了
			if i+1 == len(s) {
				ret = append(ret, w)
			}
			for _, sfx := range sfxs {
				ret = append(ret, w+" "+sfx)
			}
		}
		mem[start] = ret
		return ret
	}
	return dfs(0)
}

// @lc code=end