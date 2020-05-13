import "sort"

/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */

// @lc code=start

// 看懂别人的答案有多豁然开朗，默写出来自己却 AC 不了的时候就有多痛不欲生
func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	var ans []int
	dp, pre := make([]int, len(nums)), make([]int, len(nums))
	var max, end int
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	for i := 0; i < len(nums); i++ {
		pre[i] = -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] <= dp[j] {
				dp[i] = dp[j] + 1
				pre[i] = j
			}
		}
		if dp[i] > max {
			max = dp[i]
			end = i
		}
	}
	for i := end; i != -1; i = pre[i] {
		ans = append(ans, nums[i])
	}
	return ans
}

// @lc code=end

