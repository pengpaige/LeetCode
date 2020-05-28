/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
import "strconv"

//
// 题目表达不够清晰，输入的字符串是完整的 IP （如果符合长度要求的话）
// 遇到 0 不能丢弃，必须全部用上，所以 0 如果在开头的话必须作为单独一段
func restoreIpAddresses(s string) []string {
	var ans []string
	var curr string         // current right ip segments in string
	var start, segCount int // curr start in s
	bt93(start, segCount, s, curr, &ans)
	return ans
}

func bt93(s, c int, str, curr string, ans *[]string) {
	if c > 4 {
		return
	}
	if c == 4 && s == len(str) {
		*ans = append(*ans, curr)
		return
	}

	for i := 0; i < 3; i++ {
		if s+i+1 > len(str) {
			break
		}

		seg := string(str[s : s+i+1])
		num, _ := strconv.Atoi(seg)
		if i == 2 && num > 255 {
			continue
		}

		next := ""
		if curr != "" {
			next += curr + "." + seg
		} else {
			next += curr + seg
		}
		bt93(s+i+1, c+1, str, next, ans)
		if i == 0 && str[s] == '0' {
			break
		}
	}
}

// @lc code=end

