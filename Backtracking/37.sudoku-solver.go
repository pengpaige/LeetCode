/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte)  {
    isValidSudoku(board)
}

func isValidSudoku(board [][]byte) bool {
	var dfs func() bool
	dfs = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for n := '1'; n <= '9'; n++ {
                    if !valid(board, i, j, byte(n)) {
						continue
					}
                    board[i][j] = byte(n)
					if dfs() {
						return true
					}
					// 这里有些思维量
					// 只有 dfs 到最后一个位置(也就是 board 右下角)时 dfs = true
					// 才能证明整体的数字编排有效, 并且才能逐级返回 true
					// 否则最后一次返回 false 的话
					// 所有的 dfs 层次都会返回 false, 只能尝试其他分支
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
	return dfs()
}

// 需要注意一点, 数独里要求九宫格数字出现 1-9 
// 这里的 "九宫格" 并不是每次以当前数字为中心
// 而是以数独的 9x9 均分成 9 个九宫格
// 每个九宫格位置固定且不重叠
func valid(board [][]byte, row, col int, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == n {
			return false
		}
		if board[row][i] == n {
			return false
		}
        // 行用 i/3 列用 i%3 是为了确保能取到九宫格里的每个数字
		// 如果都用 i/3 的话那么每次横纵坐标都相同, 只能取到对角线
		// 虽然除和取余的范围都相同, 但是取余相当于在除法固定结果(行号)的同时
		// 在每行都取一便 0, 1, 2
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == n {
			return false
		}
	}
	return true
}
// @lc code=end

