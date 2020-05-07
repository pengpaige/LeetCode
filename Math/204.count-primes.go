/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func countPrimes(n int) int {
    var ans int
    for i := 1; i < n; i++ {
        if isPrime(i) {
            ans++
        }
    }
    return ans
}


func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    // n 的约数一定小于 n/2 
    // 所以直接把下面的循环结束条件设置成 n/2 是第一阶段优化
    // 但是如果 p*q = n 成立
    // 那么一定存在一个 小于等于 根号 n 的数字能被 n 整除
    // 所以是否为素数就转化为要判断小于 根号 n 的数字能不能被 n 整除
    for i := 2; i*i <= n; i++ {
        if n%i  == 0 {
            return false
        }
    }
    return true
}
// @lc code=end

