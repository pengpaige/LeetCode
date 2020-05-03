// 无聊的 O(N) 方式
func isPowerOfTwoBruteForce(n int) bool {
    if n <= 0 {
        return false
    }
    
    for n != 0 && n & 1 == 0 {
        n >>= 1
    }
    if n >>= 1; n == 0 {
        return true
    }
    return false
}


// 比较 trick 的解法
// 利用 n&(-1*n) 得到最右侧的 1 的特性
// 如果 n 是 2 的幂，那么 n 的二进制表示里肯定只有一个 1
// 也就是最右侧的 1 保留其他全部变为 0 之后应该依然等于 n
func isPowerOfTwo(n int) bool {
    if n <= 0 {
        return false
    }
    return n&(-1*n) == n
}