/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mem := make(map[int]int, len(nums))
    for index, n := range nums {
		if _, ok := mem[target-n]; ok {
			return []int{mem[target-n], index}
		}
		mem[n] = index
	}
	return nil
}
// @lc code=end