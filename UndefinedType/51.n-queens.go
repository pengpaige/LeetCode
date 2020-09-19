/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	ans := make([][]int, 0)
	// 下面是回溯记忆操作, 用于剪支
	// 因为每次回溯都按行 row 从上到下进行, 因此同一行的皇后互斥已经在回溯流程被排除了
	// 只需要考虑列和两种对角线
	col := make(map[int]struct{}, 0)	// 代表列中是否已存在皇后
	pie := make(map[int]struct{}, 0)	// 代表撇(右上到左下)是否已存在皇后
	na := make(map[int]struct{}, 0)		// 代表捺(左上到右下)是否已存在皇后
	queens := make([]int, 0)
	// i 表示搜索进行到第 i 行
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			// 这里必须对 queens 进行深拷贝, 否则后续回溯会影响 ans 中的结果
			tmp := make([]int, len(queens))
			copy(tmp, queens)
			ans = append(ans, tmp)
		}
		for j := 0; j < n; j++ {
			_, a := col[j]; _, b := pie[i+j]; _, c := na[i-j]
			// 剪支
			if !a && !b && !c {
				col[j], pie[i+j], na[i-j] = struct{}{}, struct{}{}, struct{}{}
				queens = append(queens, j)
				dfs(i+1)
				delete(col, j); delete(pie, i+j); delete(na, i-j)
				// 因为底层数据没有变, 所以必须删除本次添加的元素, 后续回溯才能正常使用 queens
				queens = queens[:len(queens)-1]
			}
		}
	}
	dfs(0)
	return genMetric(ans, n)
}

func genMetric(ans [][]int, n int) [][]string {
	ret := make([][]string, 0)
	s := ""
	for i := 0; i < n; i++ {
		s += "."
	}
	for _, queens := range ans {
		metric := make([]string, 0)
		for _, q := range queens {
			l := []byte(s)
			l[q] = 'Q'
			metric = append(metric, string(l))
		}
		ret = append(ret, metric)
	}
	return ret
}
// @lc code=end

