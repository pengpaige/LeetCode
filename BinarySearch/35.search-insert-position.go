/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start

// 二分查找(binary search) 就是想办法理清楚 mid l r 的加减和除 2 到底意味着什么
func searchInsert(nums []int, target int) int {
	// 下面这些 corner case 和最后返回 r+1 效果完全一样
	// if target > nums[len(nums)-1] {
	// 	return len(nums)
	// }
	// if target < nums[0] {
	// 	return 0
	// }
	mid, l, r := 0, 0, len(nums)-1
	for l <= r {
		mid = (l+r)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid-1
		} else if nums[mid] < target {
			l = mid+1
		}
	}
	// 这里返回 l 或者 r+1 都可以
	// 因为找不到 target 的时候返回的下标是看做将该下标以及以后的元素整体右移一位
	// 然后在将该下标作为需要返回的插入位置下标
	return l
}
// @lc code=end

