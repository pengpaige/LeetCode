import "sort"

/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */

// @lc code=start

// 这题我只花了五分钟
// 就决定去看答案了
func minEatingSpeed(piles []int, H int) int {
	// 通过排序找到 piles 中的最大值
	sort.Slice(piles, func(a, b int) bool {
		return piles[a] < piles[b]
	})
	maxp := piles[len(piles)-1]
	mid, low, high := 0, 0, maxp
	// 一定存在一个最小的 K
	// 使得当每小时吃香蕉的数量大于 K 时都能在 H 小时内吃完
	// 当然这里面隐藏了一个条件，就是 H >= len(piles)
	// 因为题中要求一小时只能吃一个 pile
	for low <= high {
		mid = (low + high) / 2
		if possible(piles, H, mid) {
			if mid == 1 {
				return 1
			} else if mid-1 >= 1 && !possible(piles, H, mid-1) {
				return mid
			}
			high = mid - 1
		} else {
			if mid+1 <= maxp && possible(piles, H, mid+1) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return -1
}

// 返回给定的 K 能否在 H 时间内吃完所有的 piles
func possible(piles []int, H, K int) bool {
	var times int
	for _, p := range piles {
		times += (p-1)/K + 1
	}
	return times <= H
}

// @lc code=end

