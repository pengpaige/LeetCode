func lengthOfLastWord(s string) int {
    if len(s) == 0 {
        return 0
    }
    end := len(s)-1
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            break
        }
        end--
    }
    start := end
    for i := end; i >= 0; i-- {
        if s[i] == ' ' {
            break
        }
        start--
    }
    return end-start
}
