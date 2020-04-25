/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
    for i := len(digits)-1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] < 10 {
			return digits
		}
		digits[i] -= 10
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
// @lc code=end

