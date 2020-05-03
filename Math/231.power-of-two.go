func isPowerOfTwo(n int) bool {
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