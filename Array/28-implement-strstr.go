func strStrBruteForce(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    h := []byte(haystack)
    n := []byte(needle)
    
    for i := 0; i < len(h)-len(n)+1; i++ {
        if h[i] == n[0] {
            j := 1
            for ; j < len(n); j++ {
                if h[i+j] != n[j] {
                    break
                }
            }
            if j == len(n) {
                return i
            }
        }
    }
    return -1
}

func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    return kmp([]byte(haystack), []byte(needle))
}

func kmp(haystack, needle []byte) int {
    next := getNext(needle)
    j := 0
    for i := 0; i < len(haystack); i++ {
        for j > 0 && needle[j] != haystack[i] {
            j = next[j-1]+1
        }
        if needle[j] == haystack[i] {
            j++
        }
        if j == len(needle) {
            return i-len(needle)+1
        }
    }
    return -1
}

func getNext(n []byte) []int {
    next := make([]int, len(n))
    next[0] = -1
    k := -1 // k 表示已完成的 next 下标
    for i := 1; i < len(n); i++ {
        for k > -1 && n[k+1] != n[i] {
            k = next[k]
        }
        if n[k+1] == n[i] {
            k++
        }
        next[i] = k
    }
    return next
}
