/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	ans, l, r := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			l[i] = 1
			continue
		}
		l[i] = l[i-1] * nums[i-1]
	}

	for i := len(nums - 1); i >= 0; i-- {
		if i == len(nums)-1 {
			l[i] = 1
			continue
		}
		l[i] = l[i+1] * nums[i+1]
	}

	for i := range nums {
		ans[i] = l[i] * r[i]
	}
	return ans
}

// O(N*2) 解法
// Your runtime beats 9.09 % of golang submissions
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		curr := 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			curr *= nums[j]
		}
		ans[i] = curr
	}
	return ans
}

// @lc code=end

