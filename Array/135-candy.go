func candy(ratings []int) int {
    c := make([]int, len(ratings))
    for i := 0; i < len(ratings); i++ {
        c[i] = 1
    }
    flag := true
    for flag {
        flag = false
        for i, _ := range ratings {
            if i < len(ratings)-1 && ratings[i] > ratings[i+1] &&c[i] <= c[i+1] {
                flag = true
                c[i] = c[i+1] + 1
            }
            if i > 0 && ratings[i] > ratings[i-1] && c[i] <= c[i-1] {
                flag = true
                c[i] = c[i-1] + 1
            }
        }
    }
    var ans int
        for i := 0; i < len(ratings); i++ {
        ans += c[i]
    }
    return ans
}
