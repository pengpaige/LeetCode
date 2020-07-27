/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome_map(s string) int {
	var pairCount int
	l := len(s)
	mem := make(map[byte]int, 0)
	for i := 0; i < l; i++ {
		mem[s[i]]++
	}
	for _, count := range mem {
		dc := count / 2
		if dc < 0 {
			continue
		}
		pairCount += dc
	}
	if pairCount*2 == l {
		return 2 * pairCount
	}
	return 2*pairCount + 1
}

// @lc code=end

