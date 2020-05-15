/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	sz := len(nums)
	ans := &[][]int{}
	// cmb := &[]int{}
	cmb := make([]int, 0, sz)

	for i := 0; i <= sz; i++ {
		bt(i, 0, cmb, ans, nums)
	}
	return *ans
}

// func bt(l, s int, cmb *[]int, ans *[][]int, nums []int) {
// 	sz := len(nums)
// 	if len(*cmb) == l {
// 		// cmb 直接用指针的话，会出现之前追加到 ans 里的变量也被覆盖成新的 cmb 的问题
// 		// 怀疑这里最终 append 到 ans 里的是 cmb 的指针？
// 		*ans = append(*ans, *cmb)
// 		return
// 	}

// 	// 这个位置 i < sz 不能优化成 i <= sz-l
// 	// 因为整个回溯过程中只有每个 s == 0 的位置才需要满足这个条件
// 	// 后面的递归分支里需要取到 i > sz-l
// 	for i := s; i < sz; i++ {
// 		*cmb = append(*cmb, nums[i])
// 		bt(l, i+1, cmb, ans, nums)
// 		*cmb = (*cmb)[:len(*cmb)-1]
// 	}
// }

func bt(l, s int, cmb []int, ans *[][]int, nums []int) {
	sz := len(nums)
	if len(cmb) == l {
		tmp := make([]int, len(cmb), len(cmb))
		copy(tmp, cmb)
		*ans = append(*ans, tmp)
		return
	}

	// 这个位置 i < sz 不能优化成 i <= sz-l
	// 因为整个回溯过程中只有每个 s == 0 的位置才需要满足这个条件
	// 后面的递归分支里需要取到 i > sz-l
	for i := s; i < sz; i++ {
		cmb = append(cmb, nums[i])
		bt(l, i+1, cmb, ans, nums)
		cmb = cmb[:len(cmb)-1]
	}
}

// @lc code=end

