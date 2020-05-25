/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */

// @lc code=start

// 全局缓存这种把自行车改装成电动车的方式只能帮你 beats 25%, 想拿双百还是得想到更好的办法而不是优化烂算法
var gm = make(map[int][][]int, 0)

func getPermutation(n int, k int) string {
	var pmt [][]int
	if _, ok := gm[n]; ok {
		pmt = gm[n]
	} else {
		pmt = permute(n)
		gm[n] = pmt
	}
	kth := pmt[k-1]
	var ans string
	for _, n := range kth {
		ans += string('0' + n)
	}
	return ans
}

func permute(l int) [][]int {
	var ans [][]int
	pmt, m := make([]int, 0, l), make(map[int]struct{}, l)
	bt(l, pmt, &ans, m)
	return ans
}

func bt(l int, pmt []int, ans *[][]int, m map[int]struct{}) {
	if len(pmt) == l {
		tmp := make([]int, l, l)
		copy(tmp, pmt)
		*ans = append(*ans, tmp)
		return
	}
	for i := 1; i <= l; i++ {
		if _, ok := m[i]; ok {
			continue
		}
		m[i] = struct{}{}
		pmt = append(pmt, i)
		bt(l, pmt, ans, m)
		delete(m, i)
		pmt = pmt[:len(pmt)-1]
	}
}

// @lc code=end

