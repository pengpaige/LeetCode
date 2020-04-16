/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start

import "strings"

func wordPattern(pattern string, str string) bool {
	l := len(pattern)
	p2s, s2p := make(map[byte]string, l), make(map[string]byte, l)
	words := strings.Split(str, " ")
	if len(words) != l {
		return false
	}
	for i := 0; i < l; i++ {
		if p2s[pattern[i]] == "" && s2p[words[i]] == 0 {
			p2s[pattern[i]] = words[i]
			s2p[words[i]] = pattern[i]
		} else if p2s[pattern[i]] != words[i] ||
			s2p[words[i]] != pattern[i] {
			return false
		}
	}
	return true
}
// @lc code=end

