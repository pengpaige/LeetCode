func canConstruct(ransomNote string, magazine string) bool {
    if len(ransomNote) > len(magazine) {
        return false
    }
    zineMap := make(map[rune]int, 0)
    for _, b := range magazine {
        zineMap[b]++
    }
    for _, b := range ransomNote {
        if count, ok := zineMap[b]; !ok || count < 1 {
            return false
        } else {
            zineMap[b]--
        }
    }
    return true
}
