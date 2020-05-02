/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

// 313/313 cases passed (800 ms)
// Your runtime beats 16.2 % of golang submissions
// Your memory usage beats 100 % of golang submissions (29.5 MB)

// 直接暴力三重也应该是可以解的，但是那样太 low
// 结合 two sum 的经验，可以固定一个数字 a，然后求剩余数字满足和为 -a
func threeSum(nums []int) [][]int {
	var ans [][]int
	// 去重，排除输入中相同的数字
	dpNumMap := make(map[int]struct{}, 0)
	// 去重，排除相同的三元组
	dpTplMap := make(map[[3]int]struct{}, 0)
    for i, n := range nums {
		if _, ok := dpNumMap[n]; ok {
			continue
		}
		dpNumMap[n] = struct{}{}
		rm := make([]int, 0)
		if ts := twoSum(append(append(rm, nums[0:i]...), nums[i+1:len(nums)]...), -1*n); ts != nil {
			tpls := make([][]int, 0)
			for i := 0; i < len(ts); i++ {
				ts[i] = append(ts[i], n)
			}
			for i := 0; i < len(ts); i++ {
				a := s2a(ts[i])
				if _, ok := dpTplMap[a]; ok {
					continue
				}
				dpTplMap[a] = struct{}{}
				tpls = append(tpls, ts[i])
			}
			ans = append(ans, tpls...)
		}
	}
	return ans
}


func twoSum(nums []int, target int) [][]int {
	var ret [][]int
	mem := make(map[int]struct{}, 0)
	for _, n := range nums {
		if _, ok := mem[target-n]; ok {
			ret = append(ret, []int{target-n, n})
			// 去重，兼容和为 target 的数有两对以上的情况
			delete(mem, target-n)
			delete(mem, n)
		} else {
			mem[n] = struct{}{}
		}
	}
	return ret
}


func s2a(s []int) [3]int {
	var ret [3]int
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	for idx, i := range s {
		ret[idx] = i
	}
	return ret
}
// @lc code=end

