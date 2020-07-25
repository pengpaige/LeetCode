/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	var ans int
	if len(grid) == 0 {
		return ans
	}
	nc, nr := len(grid), len(grid[0])
	var dfs func(c, r int)
	// 通过深度优先遍历，将同一个岛全部置 0
	dfs = func(c, r int) {
		grid[c][r] = '0'
		if c > 0 && grid[c-1][r] == '1' {
			dfs(c-1, r)
		}
		if c < nc-1 && grid[c+1][r] == '1' {
			dfs(c+1, r)
		}
		if r > 0 && grid[c][r-1] == '1' {
			dfs(c, r-1)
		}
		if r < nr-1 && grid[c][r+1] == '1' {
			dfs(c, r+1)
		}
		return
	}

	for c := 0; c < nc; c++ {
		for r := 0; r < nr; r++ {
			if grid[c][r] == '1' {
				ans++
				dfs(c, r)
			}
		}
	}
	return ans
}

// @lc code=end

