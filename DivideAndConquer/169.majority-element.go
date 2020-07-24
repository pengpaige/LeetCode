/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	ans, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == ans {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans = nums[i]
			count = 1
		}
	}
	return ans
}

// @lc code=end

