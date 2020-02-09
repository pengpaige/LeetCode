// 先统计相同位置相同的个数，即为 bull
// 在统计排除上面 bull 之后的相同数字个数，这里有一个 trick 的数学结论
// 排除 bull 之后，对于 secret 和 guess 都包含的数字取两边包含该数字次数的较小值（交集）就是 cow
func getHint(secret string, guess string) string {
    A, B := 0, 0
    ma := make(map[byte]int)
    mb := make(map[byte]int)
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            A++
        } else {
            ma[secret[i]]++
            mb[guess[i]]++
        }
    }
    ints := []byte("0123456789")
    for _, n := range ints {
        min := ma[n]
        if mb[n] < min {
            min = mb[n]
        }
        B += min
    }
    return fmt.Sprintf("%dA%dB", A, B)
}
