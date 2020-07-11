/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var ans float64
	tl := len(nums1) + len(nums2)
	if tl == 0 {
		return 0.0
	}
	mg := merge(nums1, nums2)
	if tl%2 == 0 {
		ans = float64(mg[tl/2]+mg[tl/2-1]) / 2
	} else {
		ans = float64(mg[tl/2])
	}
	return ans
}

func merge(nums1, nums2 []int) []int {
	var ret []int
	len1, len2 := len(nums1), len(nums2)
	if len1 == 0 {
		return nums2
	} else if len2 == 0 {
		return nums1
	}
	p1, p2 := 0, 0
	for p1 < len1 || p2 < len2 {
		if p1 < len1 && p2 < len2 {
			if nums1[p1] < nums2[p2] {
				ret = append(ret, nums1[p1])
				p1++
			} else {
				ret = append(ret, nums2[p2])
				p2++
			}
		} else if p1 < len1 {
			ret = append(ret, nums1[p1])
			p1++
		} else if p2 < len2 {
			ret = append(ret, nums2[p2])
			p2++
		} else {
			break
		}
	}
	return ret
}

// @lc code=end

