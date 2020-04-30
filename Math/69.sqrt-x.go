/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
// Brute Force
// 竟然不超时
func mySqrtBruteForce(x int) int {
	i := 0
	for (i+1)*(i+1) <= x {
		i++
	}
	return i
}


// Binary Search
// 用二分法实现的话，类似于二分查找排序数组中最后一个小于目标值的位置
func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := (l+r)/2
		mp2 := mid*mid 
		if mp2 == x {
			return mid
		} else if mp2 < x {
			if (mid+1)*(mid+1) > x{
				return mid
			}
			l = mid+1
		} else if mp2 > x {
			if (mid-1)*(mid-1) < x{
				return  mid-1
			}
			r = mid-1
		}
	}
	return l
}
// @lc code=end

