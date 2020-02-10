func generate(numRows int) [][]int {
    ans := make([][]int, numRows)
    if numRows == 0 {
        return ans
    }
    ans[0] = []int{1}
    for i := 1; i < numRows; i++ {
        s := make([]int, i+1)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                s[j] = 1
            } else {
                s[j] = ans[i-1][j] + ans[i-1][j-1]
            }
        }
        ans[i] = s
    }
    return ans
}
