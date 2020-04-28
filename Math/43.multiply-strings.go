/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
import(
    "strconv"
)

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    ans := make([]byte, 0)
    levels := make([][]int, len(num1))
    
    for i := len(num1)-1; i >= 0; i-- {
        dp, carry := 0, 0
        di := int(num1[i] - '0')
        l := make([]int, 0)
        for j := len(num2)-1; j >= 0; j-- {
            dj := int(num2[j] - '0')
            dp = di * dj + carry
            carry = dp / 10
            dp %= 10
            l = append([]int{dp}, l...)
        }
        if carry > 0 {
            l = append([]int{carry}, l...)
        }
        levels[len(num1) - 1 - i] = l
    }
    
    for i := 1; i < len(levels); i++ {
        zeros := make([]int, i)
        levels[i] = append(levels[i], zeros...)
    }
    
    // return itoa(levels[0]) + " " + itoa(levels[1]) + " " + itoa(levels[2])
    
    carry := 0
    for i := 0; i < len(levels[len(levels)-1]); i++ {
        column := 0
        for j := 0; j < len(levels); j++ {
            if i > len(levels[j])-1 {
                continue
            }
            column += levels[j][len(levels[j]) - 1 - i]
        }
        // return strconv.Itoa(column)
        column += carry
        carry = column / 10
        column %= 10
        ans = append([]byte{byte('0' + column)}, ans...)
        // if i == 4 {
        //     return string(ans)
        // }
    }
    if carry > 0 {
        ans = append([]byte{byte('0' + carry)}, ans...)
    }
    
    return string(ans)
}

func itoa(is []int) string {
    var ret string
    for _, i := range is {
        ret += strconv.Itoa(i)
    }

    return ret
}
// @lc code=end

