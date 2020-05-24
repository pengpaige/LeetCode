/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start

// 搞明白原理很见到，看懂算法也很简单，但是把原理翻译成算法非常致郁
// 详解看这里 https://dwz1.cc/vvOnueo
// 关联数字的字典排序问题 386.lexicographical-numbers
func nextPermutation(nums []int)  {
	sz := len(nums)
    if sz <= 1 {
		return
	}

	i, j, k := sz-2, sz-1, sz-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] {
			k-- 
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	for m, n := j, sz-1; m < n; {
		nums[m], nums[n] = nums[n], nums[m]
		m++
		n--
	}

	return
}
// @lc code=end

