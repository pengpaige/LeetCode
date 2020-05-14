/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	ans := &[][]int{}
	cmb := &[]int{}
	sz := len(nums)

	for i := 0; i <= sz; i++ {
		bt(i, 0, cmb, ans, nums)
	}
	return *ans
}

func bt(l, s int, cmb *[]int, ans *[][]int, nums []int) {
	sz := len(nums)
	if len(*cmb) == l {
		*ans = append(*ans, *cmb)
		return
	}

	// 这个位置 i < sz 不能优化成 i <= sz-l
	// 因为整个回溯过程中只有每个 s == 0 的位置才需要满足这个条件
	// 后面的递归分支里需要取到 i > sz-l
	for i := s; i < sz; i++ {
		*cmb = append(*cmb, nums[i])
		bt(l, i+1, cmb, ans, nums)
		*cmb = (*cmb)[:len(*cmb)-1]
	}
}


func TestLC78(t *testing.T) {
	ans := subsets([]int{1,2}, t)
	t.Log(ans)
}

func subsets(nums []int, t *testing.T) [][]int {
	ans := &[][]int{}
	cmb := &[]int{}
	sz := len(nums)

	for i := 0; i <= sz; i++ {
		bt(i, 0, cmb, ans, nums, t)
	}
	return *ans
}

func bt(l, s int, cmb *[]int, ans *[][]int, nums []int, t *testing.T) {
	sz := len(nums)
	if len(*cmb) == l {
		t.Logf("append l: %d, cmd: %d", l, *cmb)
		*ans = append(*ans, *cmb)
		t.Logf("after append l: %d, *ans: %d", l, *ans)
		return
	}

	// 这个位置 i < sz 不能优化成 i <= sz-l
	// 因为整个回溯过程中只有每个 s == 0 的位置才需要满足这个条件
	// 后面的递归分支里需要取到 i > sz-l
	for i := s; i < sz; i++ {
		t.Logf("l: %d, i: %d, nums[i]: %d", l, i, nums[i])
		*cmb = append(*cmb, nums[i])
		t.Log(*cmb)
		bt(l, i+1, cmb, ans, nums, t)
		*cmb = (*cmb)[:len(*cmb)-1]
	}
}
// @lc code=end

