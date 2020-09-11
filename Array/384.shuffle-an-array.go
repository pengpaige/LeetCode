import "math/rand"

/*
 * @lc app=leetcode id=384 lang=golang
 *
 * [384] Shuffle an Array
 */

// @lc code=start
// 题解 https://leetcode-cn.com/problems/shuffle-an-array/solution/xi-pai-suan-fa-shen-du-xiang-jie-by-labuladong/
type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	return Solution{original: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	l := len(this.original)
	ret := make([]int, l)
	copy(ret, this.original)
	// 洗牌算法
	for i := l - 1; i >= 0; i-- {
		// 产生 [0, i] 范围内的随机数
		r := rand.Intn(i + 1)
		ret[i], ret[r] = ret[r], ret[i]
	}
	return ret
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

