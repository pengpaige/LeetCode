func convertToTitle(n int) string {
    var ans []byte
    for n != 0 {
        n1 := (n-1) % 26
        ans = append([]byte{byte(int8('A')+int8(n1))}, ans...)
        n = (n-1) / 26
    }
    return string(ans)
}
