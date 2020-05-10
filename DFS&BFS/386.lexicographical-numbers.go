/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start


// 当你还想用字符串排序来偷懒的时候
// 评论区的大佬们随手写完一个超简洁 DFS 斩获几十个 offer


// DFS 的思维方式还是需要刻意训练
// 按顺序往结果里塞数据这种方式也要记得
func lexicalOrder(n int) []int {
	var ans []int
	dfs(n, 0, &ans)
	return ans
}


func dfs(n, currNum int, ans *[]int) {
	if currNum > n {
		return
	}
	if currNum != 0 {
		*ans = append(*ans, currNum)
	}
	for i := 0; i < 10; i++ {
		nextNum := currNum*10 + i
		if nextNum == 0 {
			continue
		}
		dfs(n, nextNum, ans)
	}
}
// @lc code=end

