/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
// 因为不想多花十分钟把示意图画出来，导致多花了一小时硬想，最后终于明白画图还是必须的。
func search(nums []int, target int) int {
	mid, l, r := 0, 0, len(nums)-1
	for l <= r {
		mid = (l + r) / 2
		if target == nums[mid] {
			return mid
		} else if nums[mid] > target {
			// target < nums[l] 意味着 l 和 target 中间隔着 rotate 之后的最大值
			// 为了避免 (l+r)/2 可能移动距离过大导致直接将 target 排除掉
			// 所以只能保守的每次对 l 自增，让 l 跨过 rotate 的最大值进入同一个区间
			// 进入同一个区间之后就可以使用普通的二分法进行搜索
			if target >= nums[l] {
				r = mid - 1
			} else {
				l++
			}
			// 逻辑和 nums[mid] > target 时类似
			// 想不明白的话就画图吧，图画出来结果非常明显
		} else if nums[mid] < target {
			if target <= nums[r] {
				l = mid + 1
			} else {
				r--
			}
		}
	}
	return -1
}

// @lc code=end

