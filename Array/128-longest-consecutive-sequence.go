func longestConsecutive(nums []int) int {
    count := 0
    numsMap := make(map[int]struct{}, 0)
    for _, num := range nums {
        numsMap[num] = struct{}{}
    }
    for _, num := range nums {
        if _, ok := numsMap[num-1]; !ok {
            currCount := 1
            currNum := num
            for inMap(numsMap, currNum+1) {
                currCount++
                currNum++
            }
            count = max(count, currCount)
        } 
    }
    return count
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func inMap(m map[int]struct{}, n int) bool {
    _, ok := m[n]
    return ok
}
