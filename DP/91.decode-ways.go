import "strconv"

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start

// dp[i] 代表截止到长度为 i 的 s 的成功解码组合数量
// 这题很像添加了条件的斐波那契额或者爬楼梯问题
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		one, _ := strconv.Atoi(string(s[i-1]))
		if inA2Z(one) && one != 0 {
			dp[i] += dp[i-1]
		}

		if i < 2 {
			continue
		}

		// 向前两个数字中有一个 0 的话就忽略 two 的可能
		if s[i-1-1] == '0' {
			// 向前两个数字中有两个 0 的话就报错，对于这题就是返回 0
			if s[i-1] == '0' {
				return 0
			}
			continue
		}

		two, _ := strconv.Atoi(string(s[i-1-1 : i]))
		if inA2Z(two) {
			dp[i] += dp[i-2]
		}

	}
	return dp[len(s)]
}

func inA2Z(i int) bool {
	if i >= 1 && i <= 26 {
		return true
	}
	return false
}

// @lc code=end

