/*
 * @lc app=leetcode id=365 lang=golang
 *
 * [365] Water and Jug Problem
 */

// @lc code=start

// 利用贝祖定理和最大公约数的方法
// 具体解释可以看这个 https://dwz1.cc/zGX35FO
// 或者这个 https://dwz1.cc/EX1McuX
// 利用数学的解法是非常快的，但我认为对于面试来说这种方式毫无意义
// 面试官出这道题肯定不是为了考你是否了解贝祖定理
// 甚至也不太可能为了考你是否会手写求最大公约数而绕这么大一个弯
// 所以我认为后面的递归或者叫 DFS 实现方式对于面试来说更有意义
func canMeasureWater_GCD(x int, y int, z int) bool {
	if x == z || y == z {
		return true
	}
	if z > x+y {
		return false
	}
	if z%gcd(x, y) == 0 {
		return true
	}
	return false
}


func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}


// 回溯的解法
var X, Y, Z int
// 全局的 mem 可能会在不同的测试例中共享，必须要定义在 canMeasureWater 函数里面
// var mem = make(map[[2]int]struct{}, 0)


func canMeasureWater(x int, y int, z int) bool {
	X, Y, Z = x, y, z
    mem := make(map[[2]int]struct{}, 0)
	return helper(0, 0, &mem)
}


func helper(currX, currY int, mem *map[[2]int]struct{}) bool {
	if currX == Z || currY == Z || currX+currY == Z {
		return true
	}
    if _, ok := (*mem)[[2]int{currX, currY}]; ok {
		return false
	}
	(*mem)[[2]int{currX, currY}] = struct{}{}
	minX2Y := min(currX, Y-currY)
	minY2X := min(X-currX, currY)
	return helper(X, currY, mem) ||
		helper(currX, Y, mem) || 
		helper(0, currY, mem) || 
		helper(currX, 0, mem) ||
		helper(currX-minX2Y, currY+minX2Y, mem) ||
		helper(currX+minY2X, currY-minY2X, mem)
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

