import(
    "reflect"
    "sort"
)

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start

// 用最笨的方法把 OJ 系统累哭
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	cmb := make([]int, 0)
	sz := len(nums)

	for i := 0; i <= sz; i++ {
		bt(nums, i, 0, &cmb, &ans)
	}
    
    for i := 0; i < len(ans); i++ {
        sortSlice(ans[i])
    }
    
    for i := 0; i < len(ans)-1; i++ {
        for j := i + 1; j < len(ans); {
            if reflect.DeepEqual(ans[i], ans[j]) {
                ans = append(ans[:j], ans[j+1:]...)
            } else {
                j++
            }
        }
    }
	return ans
}


func bt(nums []int, l, s int, cmb *[]int, ans *[][]int) {
	sz := len(nums)
	if l == len(*cmb) {
        tmp := make([]int, len(*cmb), len(*cmb))
		copy(tmp, *cmb)
		*ans = append(*ans, tmp)
		return
	}
	for i := s; i < sz; i++ {
		*cmb = append(*cmb, nums[i])
		bt(nums, l, i+1, cmb, ans)
        *cmb = (*cmb)[:len(*cmb)-1]
	}
}


func sortSlice(s []int) {
    sort.Slice(s, func(i, j int) bool {
        return s[i] < s[j]
    })
}
// @lc code=end