/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	ans, cmax, cmin := nums[0], nums[0], nums[0]
	for i, n := range nums {
		if i == 0 {
			continue
		}
		// 需要对正负数分类处理
		if n >= 0 {
			cmax = max(n, cmax*n)
			cmin = min(n, cmin*n)
		} else {
			tmp := cmax
			cmax = max(n, cmin*n)
			cmin = min(n, tmp*n)
		}
		ans = max(ans, cmax)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

