import "strings"

/*
 * @lc app=leetcode id=686 lang=golang
 *
 * [686] Repeated String Match
 */

// @lc code=start
func repeatedStringMatch(A string, B string) int {
	t, count, maxLen := "", 0, 2*len(A)+len(B)
	for len(t) < maxLen {
		if strings.Contains(t, B) {
			return count
		}
		t += A
		count++
	}
	return -1
}

// @lc code=end

