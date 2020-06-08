/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	j, curr := 0, 0
	ans := -1<<31
	for j < len(nums) {
	// 只要当前子串和为负
	// 遇到任何一个数字相加的话
	// 和都不如这个数字本身大 
	// 所以直接替换
	 if curr < 0 {
	  curr = nums[j]
		   } else {
			   curr += nums[j]
		   }
	 if curr > ans {
	  ans = curr
	 }
	 j++
	}
	return ans
   }
// @lc code=end

