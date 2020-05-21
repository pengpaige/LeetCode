/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

// @lc code=start

// 沉浸在学会回溯求组合的我被评论区优雅秀出DP的大佬们无情拍醒

// 这题要求将顺序不同的子集也计算到总数里
// 这有点像求“排列”的做做法
func combinationSum4(nums []int, target int) int {
    if len(nums) == 0 || nums == nil {
        return 0
    }
	var ans, curr int
	// 递归、回溯这种思想要了解，而且还得熟悉每种思想方法的实现特点
	// 比如这道题直接回溯肯定是超时的，如果想办法加记忆的话
	// 就不能直接用前几道组合题目那种出参的方式
	// 因为使用出参来做总数的话，必然会在递归里判断等于 target 然后做自增
	// 但是使用记忆的方式“通过底层返回的当前总数再累加”比较方便
	m := make(map[int]int, target)
	ans = bt(target, curr, nums, m)
	return ans
}

func bt(target, curr int, nums []int, m map[int]int) int {
	if curr > target {
		return 0
	}
	if curr == target {
		return 1
	}
	if _, ok := m[curr]; ok {
		return m[curr]
	}
	var ret int
	for i := 0; i < len(nums); i++ {
		// 使用 tmp 的方式和前几道组合题的加入 slice 再退出 slice 道理是一样的
		tmp := curr + nums[i]
		ret += bt(target, tmp, nums, m)
	}
	m[curr] = ret
	return ret
}

// @lc code=end

