/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
// 排序 + 双指针
// 是的，我在抄作业。
func threeSumClosest(nums []int, target int) int {
	best := -1<<31
    update := func(new int) {
		if abs(target-best) > abs(target-new) {
			best = new
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var a, b, c int
	for a = 0; a < len(nums); a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b, c = a+1, len(nums)-1
		for b < c {
			currSum := nums[a] + nums[b] + nums[c]
			if currSum == target {
				return target
			}
			update(currSum)
			if currSum > target {
				c--
				for c > b && nums[c] == nums[c+1] {
					c--
				}
			} else {
				b++
				for b < c && nums[b] == nums[b-1] {
					b++
				}
			}
		}
	}
	return best
}

func abs(i int) int {
	if i < 0 {
		return -1*i
	}
	return i
}
// @lc code=end

