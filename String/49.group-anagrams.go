/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * convert between []byte and string is stupid 
 * but I have no idea of better ways
 */

// @lc code=start

import "sort"

func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
		return nil
	}
	m := make(map[string][]string, 0)
	for _, str := range strs {
		sb := []byte(str)
		sort.Slice(sb, func(i, j int) bool {
			return sb[i] < sb[j]
		})
		sbk := string(sb)
		m[sbk] = append(m[sbk], str)
	}

	ans := make([][]string, 0)
	for _, sbs := range m {
		ans = append(ans, sbs)
	}
	return ans
}
// @lc code=end

