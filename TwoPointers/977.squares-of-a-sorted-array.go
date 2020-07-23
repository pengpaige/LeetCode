/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares_(A []int) []int {
	var ans []int
	left, right := 0, len(A)-1
	for left <= right {
        lp, rp := abs(A[left]), abs(A[right])
		if A[left] <=0 && A[right] >=0 {
			if lp > rp {
				ans = append([]int{A[left]*A[left]}, ans...)
				left++
			} else if lp <= rp {
				ans = append([]int{A[right]*A[right]}, ans...)
				right--
            }
        } else if A[left] <=0 {
            ans = append([]int{A[left]*A[left]}, ans...)
            left++
        } else {
            ans = append([]int{A[right]*A[right]}, ans...)
            right--
        }
    }
    return ans
}

func sortedSquares(A []int) []int {
    result := make([]int, len(A))
    i, j := 0, len(A) - 1
    for p := len(A) - 1; p >= 0; p-- {
        if abs(A[i]) > abs(A[j]) {
            result[p] = A[i] * A[i]
            i++
        } else {
            result[p] = A[j] * A[j]
            j--
        }
    }
    return result
}

func abs(a int) int {
    if a < 0 {
        return -1*a
    }
    return a
}


// @lc code=end

