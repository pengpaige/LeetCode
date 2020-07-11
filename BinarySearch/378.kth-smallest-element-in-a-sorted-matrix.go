import "sort"

/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	d := len(matrix)
	s := make([]int, 0)
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			s = append(s, matrix[i][j])
		}
	}
	sort.Slice(s, func(a, b int) bool {
		return s[a] < s[b]
	})
	return s[k-1]
}

func kthSmallest_WA(matrix [][]int, k int) int {
	d := len(matrix)
	if d == 1 {
		return matrix[0][0]
	}
	n := getN(d, k)
	// return n
	last := make([]int, 0)
	idxInLast := k - count(d, n-1) - 1
	// return idxInLast
	if n <= d {
		for i := 0; i < n; i++ {
			last = append(last, matrix[i][n-1-i])
		}
	} else {
		for i := 0; i < d-(n-d); i++ {
			last = append(last, matrix[d-1-i][n-d+i])
		}
	}
	sort.Slice(last, func(a, b int) bool {
		return last[a] < last[b]
	})
	return last[idxInLast]
}

// 找到第一个满足 gauss(mid) >= k 的 mid
func getN(d, k int) int {
	mid, lo, hi := 0, 0, k
	for lo <= hi {
		mid = (lo + hi) / 2
		if count(d, mid) == k {
			return mid
		} else if count(d, mid) < k {
			if mid <= hi && count(d, mid+1) > k {
				return mid + 1
			}
			lo = mid + 1
		} else if count(d, mid) > k {
			if mid >= lo && count(d, mid-1) > k {
				return mid
			}
			hi = mid - 1
		}
	}
	return -1
}

func gauss(n int) int {
	return (n + 1) * n / 2
}

// d 是 matrix 的维度
// n 是第 n 条斜线路径
func count(d, n int) int {
	g := gauss(n)
	if n <= d {
		return g
	}
	left := d - 1 - (n - d)
	return g + gauss(d-1) - gauss(left)
}

// @lc code=end

