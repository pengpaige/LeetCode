/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		ts := numbers[l] + numbers[r]
		if ts == target {
			return []int{l+1, r+1}
		} else if ts < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
// @lc code=end

