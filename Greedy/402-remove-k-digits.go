func removeKdigits(num string, k int) string {
    // 下面的代码没办法处理去除全部 位 的情况
    if len(num) == k {
        return "0"
    }
    ans, orgK := make([]byte, 0), k
    for _, r := range num {
        d := byte(r)
        if len(ans) == 0 {
            ans = append(ans, d)
            continue
        }
        for k > 0 && len(ans) > 0 && d < ans[len(ans)-1] {
            k--
            ans = ans[:len(ans)-1]
        }
        // 这里包含 k>0 时的入栈和 k=0 后的入栈两种情况
        ans = append(ans, d)
    }
    if k > 0 {
        ans = ans[:len(num)-orgK]
    }
    // 覆盖 0 开头的测试例
    for len(ans) > 1 && ans[0] == '0' {
        ans = ans[1:]
    }
    return string(ans)
}
