import(
    "sort"
)

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	cmb := make([]int, 0)
	sz := len(nums)
    
    // 排序是为了把相同的数字放到一起，方便后面判重
    sort.Slice(nums, func(i, j int) bool {
        return s[i] < s[j]
    })

	for i := 0; i <= sz; i++ {
		bt(nums, i, 0, &cmb, &ans)
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
        if i > s && nums[i] == nums[i-1] {
            continue
        }
		*cmb = append(*cmb, nums[i])
		bt(nums, l, i+1, cmb, ans)
        *cmb = (*cmb)[:len(*cmb)-1]
	}
}
// @lc code=end