import "sort"

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start

// 沉醉在学会了组合数一般解法的喜悦当中
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	// 实际上 sz 的范围可以简单的认为是小于等于 target 
	// 即假设 candidates 中有 1 的话最多重复 target 次
	// 但是这样的话，会导致递归的层数过多而超时
	// 所以必须得把子集的最大长度缩小
	sz := target/candidates[0]
	cmb := make([]int, 0, sz)
	for l := 1; l <= sz; l++ {
		bt(candidates, cmb, l, 0, target, &ans)
	}
	return ans
}

func bt(candidates, cmb []int, l, s, target int, ans *[][]int) {
	if len(cmb) == l {
		if sum(cmb) != target {
			return
		}
		tmp := make([]int, len(cmb))
		copy(tmp, cmb)
		*ans = append(*ans, tmp)
		return
	}
	for i := s; i < len(candidates); i++ {
		// 和普通的求全部组合 78 题不同之处在与 下一个 s 从 i 开始而不是 i+1
		cmb = append(cmb, candidates[i])
		bt(candidates, cmb, l, i, target, ans)
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

