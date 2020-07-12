import "sort"

/*
 * @lc app=leetcode id=719 lang=golang
 *
 * [719] Find K-th Smallest Pair Distance
 */

// @lc code=start

// 这道题和 378.kth-smallest-element-in-a-sorted-matrix 非常类似
// 只是在求 mid 满足要求的数量时稍微有点区别
// 看来二分查找类型的题难度提升只能在边界判断的位置做文章了
func smallestDistancePair(nums []int, k int) int {
	// 先给 nums 排序，测试例要求 nums 长度最少是 2, 不用做长度判断
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	mid, lo, hi := 0, 0, nums[len(nums)-1]-nums[0]
	for lo <= hi {
		mid = (lo + hi) / 2
		c := shorterCount(nums, mid)
		if c >= k {
			if shorterCount(nums, mid-1) < k {
				return mid
			}
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return mid
}

// @guess 猜测的满足要求的数字
func shorterCount(nums []int, guess int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= guess {
				count++
			}
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// @lc code=end

