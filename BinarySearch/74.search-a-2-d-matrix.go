/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rowCount, lineCount, row := len(matrix), len(matrix[0]), 0
	
	for row < rowCount {
		if target >= matrix[row][0] && target <= matrix[row][lineCount-1] {
			break
		}
		row++
	}
	if row == rowCount {
		return false
	}

	l, r, mid := 0, lineCount-1, 0
	for l <= r {
		mid = (l+r)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			r = mid-1
		} else {
			l = mid+1
		}
	}
	return false
}
// @lc code=end

