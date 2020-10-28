/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start

// 基于二分查找的解法 时间复杂度 O(log(m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nusm2)
	// 下面的 left 和 right 值是为了兼容两个数组长度和奇偶的情况
	// m+n 是奇数时 left 和 right 相等, 代表中位数
	// m+n 是偶数时 left 和 right 相等, 代表中间两个数字
	// ps: left 和 right 代表的数字位置从 1 开始而非从 0 开始
	left, right := (m+n+1)/2, (m+n+2)/2
	if left == right {
		return getKthLeast(nums1, nums2, 0, m-1, 0, n-1, left)
	}
	return float64(getKthLeast(nums1, nums2, 0, m-1, 0, n-1, left) + getKthLeast(nums1, nums2, 0, m-1, 0, n-1, right))*0.5
}

func getKthLeast(nums1, nums2 []int, start1, end1, start2, end2, k) float64 {
	// 递归到目前为止, 剩下的两个数组长度
	len1, len2 := start1-end1+1, start2-end2+1
	// 为了简化后续判断逻辑, 调整顺序让 nums1 始终更短
	if len1 > len2 {
		return getKthLeast(nums2, nums1, start2, end2, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start+k-1]
	}

	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i, j := start1+min(len1, k/2)-1, start2+min(len2, k/2)-1

	if nums1[i] > nums2[j] {
		return getKthLeast(nums1, nums2 []int, start1, end1, j+1, end2, k-min(len2, k/2))
	}
	return getKthLeast(nums1, nums2 []int, i+1, end1, start2, end2, k-min(len1, k/2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

