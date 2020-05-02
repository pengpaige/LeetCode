/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start


// 四指针的方式
// Runtime: 8 ms, faster than 90.18% of Go online submissions for 4Sum.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for 4Sum.
func fourSum(nums []int, target int) [][]int {
	if len(nums)<4 {
		return nil
	}
	var ans [][]int
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
	for a := 0; a <= len(nums)-4; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a+1; b <= len(nums)-3; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			l, r := b+1, len(nums)-1
			for l < r {
				currSum := nums[a]+nums[b]+nums[l]+nums[r]
				if currSum == target {
					ans = append(ans, []int{nums[a], nums[b], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
				} else if currSum > target {
					r--
				} else if currSum < target {
					l++
				}
			}
		}
	}
	return ans
}
// @lc code=end

