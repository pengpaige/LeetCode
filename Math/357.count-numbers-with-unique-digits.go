/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	var ans int
    for i := 1; i < 10; i++ {
		mem := make(map[int]bool, 10)
		mem[i] = true
		ans += dfs(i, 1, n, mem)
	}
	return ans
}


func dfs(currNum, step, n int, mem map[int]bool) int {
	if step > n {
		return 0
	}
	var count int
	for i := 0; i < 10; i++ {
		if mem[i] {
			continue
		}
		mem[i] = true
		count += dfs(currNum*10+i, step+1, n, mem)
	}
	return count
}
// @lc code=end

