/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
// 类似二分的思路 复杂度降到 O(logN)

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n *= -1
	}
	pow := 1.0
	for n > 0 {
		if n%2 != 0 {
			pow *= x
		}
		x *= x
		n /= 2
	}
	return pow
}

func myPow_Recursion_DivideConquer(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n *= -1
	}
	return helper(x, n)
}

func helper(x float64, n int) float64 {
	if x == 0 || n == 0 {
		return 1.0
	}
	half := helper(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}

// 朴素循环
func myPowBruteForce(x float64, n int) float64 {
	if n == 0 || x == 1.0 {
		return 1
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}

	ans, positive := x, false
	if n > 0 {
		positive = true
	} else {
		n *= -1
	}

	for n > 1 {
		ans *= x
		n--
	}

	if !positive {
		return 1 / ans
	}
	return ans
}

// @lc code=end

