/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
func addDigits(num int) int {
    for num > 9 {
        var tmpNum int
        for i := num; i > 0; i = i/10 {
            tmpNum += i%10
        }
        num = tmpNum
    }
    return num
}
// @lc code=end

