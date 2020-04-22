func removeDuplicateLetters(s string) string {
    stack := make([]byte, 0)
    inStack := make(map[byte]struct{}, 0)
    countMap := make(map[byte]int, 0)
    for i := 0; i < len(s); i++ {
        countMap[s[i]]++
    }
    for i := 0; i < len(s); i++ {
        countMap[s[i]]--
        if _, ok := inStack[s[i]]; ok {
            continue
        }
        for len(stack) > 0 && s[i] < stack[len(stack)-1] && countMap[stack[len(stack)-1]] > 0 {
            delete(inStack, stack[len(stack)-1])
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, s[i])
        inStack[s[i]] = struct{}{}
    }
    return string(stack)
}
