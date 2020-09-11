/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	ans, q := make([]int, 0), make([]int, 0)
	for i := 0; i < l; i++ {
		for len(q) > 0 && nums[i] > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		q = append(q, nums[i])
		if i < k-1 {
			continue
		}
		ans = append(ans, q[0])
		// 这里不是每次移动窗口都会执行
		// 因为可能有的 nums[i-k+1] 因为不够大所以已经被剔除出窗口 q 了
		if q[0] == nums[i-k+1] {
			q = q[1:len(q)]
		}
	}
	return ans
}

// @lc code=end

