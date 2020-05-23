import "sort"
/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	var ans [][]int
	sz := len(nums)
	pmt, m := make([]int, 0, sz), make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	sort.Slice(nums, func (i, j int) bool {
		return nums[i] < nums[j]
	})
	bt(nums, pmt, &ans, m)
	return ans
}

func bt(nums, pmt []int, ans *[][]int, m map[int]int) {
	if len(pmt) == len(nums) {
		tmp := make([]int, len(nums), len(nums))
		copy(tmp, pmt)
		*ans = append(*ans, tmp)
		return
	}
	for i, n := range nums {
		// 同一个层次中如果有两个相同的数字
		// 而且同一个层次意味着已经选定了相同的前缀
		// 这样的话相同数字中的前一个选中时产生的排列肯定是可以包含后面数字选中时的排列的
		// 因此这里用这种方式排除了相同数字造成的重复排列问题
		// 这样做达到的效果其实就是把相同的数字看做是没有顺序的
		// 即 对相同数字按照组合处理 这就跟前面的组合题目中输入包含相同数字并且结果中也包含相同数字的情况相同
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		if m[n] < 1 {
			continue
		}
		m[n]--
		pmt = append(pmt, n)
		bt(nums, pmt, ans, m)
		pmt = pmt[:len(pmt)-1]
		m[n]++
	}
}

// @lc code=end

