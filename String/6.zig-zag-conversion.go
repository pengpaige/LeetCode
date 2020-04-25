/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
    if numRows < 2 {
        return s
    }
    rows := make([][]byte, numRows)
    // for i := 0; i < numRows; i++ {
    //     rows[i] = make([]byte, 0)
    // }
    flag, r := -1, 0
    for i := 0; i < len(s); i++ {
        if r == 0 || r == numRows - 1 {
            flag *= -1
        }
        rows[r] = append(rows[r], s[i])
        r += flag
    }
    var ans string
    for _, row := range rows {
        ans += string(row)
    }
    return ans
}
// @lc code=end

