func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    for i := 0; i < len(strs[0]); i++ {
        for j := 1; j < len(strs); j++ {
            if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
                return string(strs[0][:i])
            }
        }
    }
    return strs[0]
}


func longestCommonPrefixBruteForce(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    prefixLen := 0
    bstrs := make([][]byte, 0)
    for _, str := range strs {
        if len(str) == 0 {
            return ""
        }
        bstrs = append(bstrs, []byte(str))
    }
    curByte := bstrs[0][0]
    miss := false
    for {
        if prefixLen > 0 && len(bstrs[0]) - 1 < prefixLen {
            prefixLen--
            break
        }
        curByte = bstrs[0][prefixLen]
        i := 0
        for ; i < len(bstrs); i++ {
            if prefixLen > len(bstrs[i])-1 || bstrs[i][prefixLen] != curByte {
                miss = true
                break
            }
        }
        if miss {
            prefixLen--
            break
        } else {
            prefixLen++
        }
    }
    ans := []byte{}
    for i := 0; i <= prefixLen; i++ {
        ans = append(ans, bstrs[0][i])
    }
    return string(ans)
}


