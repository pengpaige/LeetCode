/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	// 找不到直接返回 -1
	left, right := -1, -1
	mid, l, r := 0, 0, len(nums)-1

	// 找最右
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid] < nums[mid+1] {
				right = mid
				break
			} else if nums[mid] == nums[mid+1] {
				l = mid + 1
			}
		}
	}

	// 找最左
	l, r = 0, len(nums)-1
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			if mid == 0 || nums[mid] > nums[mid-1] {
				left = mid
				break
			} else if nums[mid] == nums[mid-1] {
				r = mid - 1
			}
		}
	}

	return []int{left, right}
}

// @lc code=end

