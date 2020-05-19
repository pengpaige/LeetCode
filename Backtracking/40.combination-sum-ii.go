import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start

// 本题和 39 的区别就是 39 输入没有重复但是子集中可以重复使用输入中的数字
// 但本题输入可能有重复 输出要求不能重复使用输入中的数字
// 当然两道题目都要结果中的求子集整体不能重复
func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	// 减少子集长度 也就是递归的层数 从而防止因代码超时导致的抑郁
	sz := target / candidates[0]
	cmb := make([]int, 0, sz)
	for l := 1; l <= sz; l++ {
		bt(candidates, cmb, l, 0, target, &ans)
	}
	return ans
}

func bt(candidates, cmb []int, l, s, target int, ans *[][]int) {
	if l == len(cmb) {
		if sum(cmb) != target {
			return
		}
		tmp := make([]int, len(cmb), len(cmb))
		copy(tmp, cmb)
		*ans = append(*ans, tmp)
		return
	}
	for i := s; i < len(candidates); i++ {
		// 回溯中同一层的数字不能相同
		// 这样可以保证输入中本来就重复的可以多次选中
		// 但是不会因为本身的重复数字导致子集前缀相同
		if i > s && candidates[i] == candidates[i-1] {
			continue
		}
		cmb = append(cmb, candidates[i])
		bt(candidates, cmb, l, i+1, target, ans)
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

