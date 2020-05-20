/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var ans [][]int
	cmb := make([]int, 0, k)
	bt(digits, cmb, k, 0, n, &ans)
	return ans
}

func bt(digits, cmb []int, l, s, target int, ans *[][]int) {
	if len(cmb) == l {
		if target != sum(cmb) {
			return
		}
		tmp := make([]int, l)
		copy(tmp, cmb)
		*ans = append(*ans, tmp)
		return
	}
	for i := s; i < len(digits); i++ {
		cmb = append(cmb, digits[i])
		bt(digits, cmb, l, i+1, target, ans)
		cmb = cmb[:len(cmb)-1]
	}
}

func sum(s []int) int {
	var sum int
	for _, n := range s {
		sum += n
	}
	return sum
}

// @lc code=end

