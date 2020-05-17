/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var ans [][]int
	e := make([]int, 0, k)
	bt(n, k, 1, e, &ans)
	return ans
}


func bt(n, k, s int, e []int, ans *[][]int) {
	if len(e) == k {
		tmp := make([]int, k, k)
		copy(tmp, e)
		*ans = append(*ans, tmp)
	}
	for i := s; i <= n; i++ {
		e = append(e, i)
		bt(n, k, i+1, e, ans)
		e = e[:len(e)-1]
	}
}
// @lc code=end

