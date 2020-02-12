func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    if rowIndex == 1 {
        return []int{1,1}
    }
    // rowIndex >= 2
    var pre, curr []int
    pre = []int{1, 1}
    for i := 2; i <= rowIndex; i++ {
        curr = make([]int, i+1)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                curr[j] = 1
            } else {
                curr[j] = pre[j] + pre[j-1]
            }
        }
        pre = curr 
    }
    return curr
}
