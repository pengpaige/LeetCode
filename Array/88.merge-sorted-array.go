/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	p1, p2, p := m-1, n-1, m+n-1
	for p2 >= 0 {
		if p1 >= 0 && p2 >= 0 {
			if nums1[p1] >= nums2[p2] {
				nums1[p] = nums1[p1]
				p1--
			} else {
				nums1[p] = nums2[p2]
				p2--
			}
		} else if p2 >= 0 { // 只要 nums2 数组全放到 nums1 中就可以结束
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	return
}

// @lc code=end

