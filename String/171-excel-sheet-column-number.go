import "math"

func titleToNumber(s string) int {
    var ans float64
    sb := []byte(s)
    for i := len(sb)-1; i >= 0; i-- {
        ans += math.Pow(26, float64(len(s)-i-1))*float64(sb[i]-'A'+1)
    }
    return int(ans)
}
