// 今天封装接口实在太累了，直接抄个答案 https://dwz1.cc/HBcRhzn

func intToRoman(num int) string {
    nums := []int{     1000, 900, 500, 400,  100, 90,   50,  40,   10,   9,    5,   4,    1}
    romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    ans, index := "", 0
    for index < 13 {
        for num >= nums[index] {
            ans += romans[index]
            num -= nums[index]
        }
        index++
    }
    return ans
}