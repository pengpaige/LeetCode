/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */

// @lc code=start


// 搞懂递归的难度仅次于搞懂女朋友的脑回路

// 回溯或者 DFS 很多时候很难统一成一个入口
// 这时候尝试遍历初始条件将 “第一层” DFS 写成多个入口比较容易理解
func countNumbersWithUniqueDigits(n int) int {
	var ans int
    mem := make(map[int]bool, 0)
    for i := 1; i < 10; i++ {
		dfs(i, 1, n, &mem, &ans)
	}
    ans++
	return ans
}


func dfs(d, step, n int, mem *map[int]bool, count *int) {
	if step > n {
		return
	}
    if (*mem)[d] {
        return
    }
    (*mem)[d] = true
    (*count)++
	for i := 0; i < 10; i++ {
		dfs(i, step+1, n, mem, count)
	}
	// 回溯的关键
    (*mem)[d] = false
}
// @lc code=end

