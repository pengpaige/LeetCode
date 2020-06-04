/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start

// DP 和贪心算法之间可能存在某种薛定谔距离

// 状态转移方程：dp[i] = dp[i-j*j] + 1
// 其中 dp[i] 数字 i 的 Perfect Square, 以下简称 PS
// 首先，数字 i 最多可以被分解成 i 个 1 的 平方，但这不一定是 PS
// 假如 j*j < i 并且没有一个更大的 j 满足 j*j < i
// 不难理解（不难理解的意思就是好像有道理但是又不会证明, 为了防止你叫我证明给你看，就故意显出一种「这还不是显而易见」的样子，只要你问为什么就是你太笨）dp[i-j*j] + 1 = dp[i]
// 上面的等式，数字 「i 的 PS」 和 「i 减去一个数字 j 的平方 j*j 的 PS」 相差 1
// 对，就是差了 j*j 这个数字，这看起来很像废话，但假如我们找到最大的小于 i 的 j*j 的话
// dp[i-j*j] + 1 = dp[i] 这个等式就是状态转移方程
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

