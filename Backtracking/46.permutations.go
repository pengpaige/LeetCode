/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	var ans [][]int
	sz := len(nums)
	pmt, m := make([]int, 0, sz), make(map[int]struct{}, sz)
	bt(sz, pmt, nums, &ans, m)
	return ans
}

func bt(l int, pmt, nums []int, ans *[][]int, m map[int]struct{}) {
	if len(pmt) == l {
		tmp := make([]int, l, l)
		copy(tmp, pmt)
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = struct{}{}
		pmt = append(pmt, nums[i])
		bt(l, pmt, nums, ans, m)
		delete(m, nums[i])
		pmt = pmt[:len(pmt)-1]
	}
}

// @lc code=end

