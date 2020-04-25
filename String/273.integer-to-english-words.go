/*
 * @lc app=leetcode id=273 lang=golang
 *
 * [273] Integer to English Words
 */

// @lc code=start
import (
    "strings"
    
)

func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    return strings.Join(helper(num), " ")
}

func helper(num int) []string {
    s19 := "One Two Three Four Five Six Seven Eight Nine Ten " +
            "Eleven Twelve Thirteen Fourteen Fifteen Sixteen Seventeen Eighteen Nineteen"
    s10s := "Twenty Thirty Forty Fifty Sixty Seventy Eighty Ninety"
    s1000s := "Thousand Million Billion"
    to19, tens, grands := strings.Split(s19, " "), strings.Split(s10s, " "), strings.Split(s1000s, " ")
    handred := "Hundred"
	
	if num == 0 {
		return nil
	} else if num < 20 {
        return to19[num-1:num]
    } else if num < 100 {
        return append([]string{tens[num/10-2]}, helper(num%10)...)
    } else if num < 1000 {
        return append(append([]string{to19[num/100-1]}, []string{handred}...), helper(num%100)...)
    }
    for i, k := range grands {
        if num < int(math.Pow(float64(1000), float64(i+2))) {
			return append(append(helper(num/int(math.Pow(float64(1000), float64(i+1)))), []string{k}...), helper(num%int(math.Pow(float64(1000), float64(i+1))))...)
        }
    }
    return nil
}
// @lc code=end

