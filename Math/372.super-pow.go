/*
 * @lc app=leetcode id=372 lang=golang
 *
 * [372] Super Pow
 */

// @lc code=start

// 1 用递归优化高阶幂
// 2 使用公式 (a%k)*(b%k)=a*b%k 同时优化大数溢出和大数求余数（取模）
// 3 递归求高阶幂的运算通过 a^2n = (a^n)^2 优化到 O(logN)
// - 考虑这么全面的解法我是不可能想到的 https://dwz1.cc/BF9Svuu
// - 这辈子都不可能想到的
var base = 1337
func superPow(a int, b []int) int {
	if a == 0 {
		return 0
	}
	if len(b) == 0 {
		return 1
	}
	lastDigit := b[len(b)-1]
	b = b[:len(b)-1]
	ans1 := powerAndMod(a, lastDigit)
	ans2 := powerAndMod(superPow(a, b), 10)
	return ans1 * ans2 % base
}


func powerAndMod(a, b int) int {
	if b == 0 {
		return 1
	}
	if b % 2 == 0 {
		h := powerAndMod(a, b/2)
		return h * h % base
	} else {
		return (a % base) * powerAndMod(a, b-1) % base
	}
}


// func powerAndMod(a, b int) int {
// 	ret := 1
// 	for i := 0; i < b; i++ {
// 		ret = ret * a % base
// 	}
// 	return ret
// }
// @lc code=end

