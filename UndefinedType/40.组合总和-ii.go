/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	return getSumKComps(candidates, target)
}

func getSumKComps(nums []int, k int) [][]int {
	var ans [][]int

	var dfs func(l, pos int, curr []int)
	dfs = func(l, pos int, curr []int) {
		if len(curr) > l {
			return
		}
		if len(curr) == l {
			if !isTargetSum(curr, k) {
				return
			}
			temp := make([]int, len(curr), len(curr))
			copy(temp, curr)
			ans = append(ans, temp)
		}
		for i := pos; i < len(nums); i++ {
			curr = append(curr, nums[i])
			dfs(l, i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	var curr []int
	for l := 1; l <= len(nums); l++ {
		dfs(l, 0, curr)
	}

	return ans
}

func isTargetSum(curr []int, target int) bool {
	var sum int
	for _, i := range curr {
		sum += i
	}
	if sum == target {
		return true
	}
	return false
}

// @lc code=end

