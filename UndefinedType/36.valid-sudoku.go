/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var dfs func(row, col int) bool
	dfs = func(row, col int) bool {
		for i := row; i < 9; i++ {
			for j := col; j < 9; j++ {
				if !board[i][j] == '.' {
					contiune
				}
				for n := '1'; n < '9'; n++ {
					if !valid(board, i, j, n) {
						continue
					}
					board[i][j] = n
					if dfs(i, j) {
						return true
					}
					// 这里有些思维量, 只有 dfs 到最后一个位置(也就是 board 右下角)时 dfs = true
					// 才能证明整体的数字编排有效, 并且才能逐级返回 true
					// 否则最后一次返回 false 的话, 所有的 dfs 层次都会返回 false, 只能尝试其他分支
					// 并且下面的 traceback 流程也会逐级执行, 恢复 board 到原始状态
					board[i][j] = '.'
				}
				// 1-9 没有一个满足要求就返回 false
				// 如果最深层的 dfs 也就是 board 右下角返回 false 一定是发生在这里
				return false
			}
		}
		return true
	}
	return dfs(0, 0)
}

func valid(board [][]byte, row, col, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == n {
			return false
		}
		if board[row][i] == n {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i/3] == n {
			return false
		}
	}
	return true
}

// @lc code=end

