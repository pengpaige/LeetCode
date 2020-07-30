/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	var ans []int
	lr, lc := len(matrix), len(matrix[0])
	for start := 0; lc > 2*start && lr > 2*start; start++ {
		ans = append(ans, printInCircle(matrix, lr, lc, start)...)
	}
	return ans
}

// nums[Y][X]
func printInCircle(nums [][]int, lr, lc, start int) []int {
	var ret []int
	endX, endY := lc-1-start, lr-1-start

	for i := start; i <= endX; i++ {
		ret = append(ret, nums[start][i])
	}

	if start < endY {
		for i := start + 1; i <= endY; i++ {
			ret = append(ret, nums[i][endX])
		}
	}

	if start < endY && start < endX {
		for i := endX - 1; i >= start; i-- {
			ret = append(ret, nums[endY][i])
		}
	}

	if start < endX && start < endY-1 {
		for i := endY - 1; i > start; i-- {
			ret = append(ret, nums[i][start])
		}
	}

	return ret
}

// @lc code=end

