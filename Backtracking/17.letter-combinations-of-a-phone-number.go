/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start

// 即使拿来双百也不能骄傲呦 ^_^
// 这基本上是常规 permutation 的变型。
// 递归写起来还是不熟练，思维很容易乱。
var nl = []string{"", "",
	"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	var ans []string
    if digits == "" {
        return ans
    }
	curr := make([]byte, 0, len(digits))
	bt(0, digits, curr, &ans)
	return ans
}

func bt(id int, digits string, curr []byte, ans *[]string) {
	if id == len(digits) {
		tmp := make([]byte, len(curr), len(curr))
		copy(tmp, curr)
		*ans = append(*ans, string(tmp))
        return
	}
	digit := int(digits[id] - '0')
	lts := nl[digit]
	for i := 0; i < len(lts); i++ {
		curr = append(curr, lts[i])
		bt(id+1, digits, curr, ans)
		curr = curr[:len(curr)-1]
	}
}
// @lc code=end

