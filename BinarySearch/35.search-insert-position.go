/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target < nums[len(nums)-1] {
		return 0
	}
	mid, l, r := 0, 0, len(nums)-1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			return mid
		} if nums[mid] > target {
			r = mid-1
		} if nums[mid] < target {
			l = mid+1
		}
	}
	return r+1
}
// @lc code=end

