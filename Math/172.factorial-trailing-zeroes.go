/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start

// 几个 5 都找不到，工作肯定也找不到，放弃吧，别刷了。
// 没戏。

// 每个 2 和 5 相乘就能得到一个 10 从而在结果里增加一个 0
// 但是因为 5 肯定比 2 少，因为每个五个数字才有一个 5
// 别说这是废话，因为你没想到
// 所以只统计 5 的个数就可以了
// 另外整数除法的好处就是默认帮你舍去余数，正好商就是 5 的个数
// 别忘了 25 和 125 这种每间隔 5 的幂出现一次恶情况

func trailingZeroes(n int) int {
	var ans int
	for n > 0 {
		ans += n / 5
		n /= 5
	}
	return ans
}

// @lc code=end

