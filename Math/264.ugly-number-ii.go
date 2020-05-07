/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start

// 亘古不变的似曾相识，亘古不变的想不起来
// 596/596 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)

var uglyNumbers = []int{1}
var i2, i3, i5 = 0, 0, 0

func nthUglyNumber(n int) int {
	for len(uglyNumbers) < n {
		min := min(min(uglyNumbers[i2]*2, uglyNumbers[i3]*3), uglyNumbers[i5]*5)
		if min == uglyNumbers[i2]*2 {
			i2++
		}
		if min == uglyNumbers[i3]*3 {
			i3++
		}
		if min == uglyNumbers[i5]*5 {
			i5++
		}
		uglyNumbers = append(uglyNumbers, min)
	}
	return uglyNumbers[n-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

