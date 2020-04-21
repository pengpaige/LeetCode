import "fmt"

func countAndSay(n int) string {
    s := "1"
    for i := 1; i < n; i++ {
        cur, tmp, count := s[0], "", 0
        for j := 0; j < len(s); j++ {
            if cur == s[j] {
                count++
            } else {
                tmp = fmt.Sprintf("%s%d%c", tmp, count, cur)
                cur = s[j]
                count = 1
            }
        }
        s = fmt.Sprintf("%s%d%c", tmp, count, cur)
    }
    return s
}
