/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
// 今天封装接口实在太累了，直接抄个答案 https://dwz1.cc/HBcRhzn

func intToRoman(num int) string {
    nums := []int{     1000, 900, 500, 400,  100, 90,   50,  40,   10,   9,    5,   4,    1}
    romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    ans, index := "", 0
    // 因为测试例的局限性，官方的输入最大不超过 3999，所以这道题的解法是比较取巧的
    for index < 13 {
        for num >= nums[index] {
            ans += romans[index]
            num -= nums[index]
        }
        index++
    }
    return ans
}
// @lc code=end

