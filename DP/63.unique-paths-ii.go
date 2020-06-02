/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start

// dp[i][j] = dp[i-1][j]+dp[i][j-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	c, l := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, c)
	for i := range dp {
		dp[i] = make([]int, l)
	}

	for i := 0; i < c; i++ {
		for j := 0; j < l; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[0][0] = 1
			} else if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[c-1][l-1]
}

// @lc code=end

