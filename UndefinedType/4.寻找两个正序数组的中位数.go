/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start

// 基于二分查找的解法 时间复杂度 O(log(m+n))
// 解法来自这条题解 https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-2/
// 整体思路，找中位数等价于找第 k 小的数
// 其中 k 表示第 k 小的数字, 等价于中位数，k 从 1 开始
// 如果中位数是两个数平均值，则等价于求两个第 k 和 k+1 小的数字
// 每次求第 k 小的数字时，为两个数组设置初始位置和结束位置
// 开始时在两个数组中各取 k/2 个数字，如果两个数组第 k/2 相等
// 那么这个数字就是正在寻找的数字
// 如果不相等，较小的数字前面的 k/2 个数字必定在中位数前面
// 这时直接忽略这 k/2 个数字，同时修改 k 的值为 k-min(lenX, k/2)
// k 修改为 k-min(lenX, k/2) 而不是直接修改为 k/2 是因为
// 可能某个数组当前选中的部分长度已经不足 k/2，这时直接比较结尾数字即可
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// 下面的 left 和 right 值是为了兼容两个数组长度和奇偶的情况
	// m+n 是奇数时 left 和 right 相同, 代表中位数
	// m+n 是偶数时 left 和 right 不同, 代表中间两个数字
	// ps: left 和 right 代表的是第 k 小的数, 因此从 1 开始而非从 0 开始
	left, right := (m+n+1)/2, (m+n+2)/2
	if left == right {
		return float64(getKthLeast(nums1, nums2, 0, m-1, 0, n-1, left))
	}
	return float64(getKthLeast(nums1, nums2, 0, m-1, 0, n-1, left)+getKthLeast(nums1, nums2, 0, m-1, 0, n-1, right)) * 0.5
}

func getKthLeast(nums1, nums2 []int, start1, end1, start2, end2, k int) int {
	// 递归到目前为止, 剩下的两个数组长度
	len1, len2 := end1-start1+1, end2-start2+1
	// 为了简化后续判断逻辑, 调整顺序让 nums1 始终更短
	if len1 > len2 {
		return getKthLeast(nums2, nums1, start2, end2, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}

	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	// 求出中间位置, 即上述题解描述的两个数组各自的第 k/2 数字
	i, j := start1+min(len1, k/2)-1, start2+min(len2, k/2)-1

	if nums1[i] > nums2[j] {
		return getKthLeast(nums1, nums2, start1, end1, j+1, end2, k-min(len2, k/2))
	}
	return getKthLeast(nums1, nums2, i+1, end1, start2, end2, k-min(len1, k/2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

