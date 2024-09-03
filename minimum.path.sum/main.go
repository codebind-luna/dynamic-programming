package main

import "math"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// move upwords
			choiceOne, choiceTwo := math.MaxInt, math.MaxInt
			if i > 0 {
				choiceOne = dp[i-1][j]
			}

			// move leftwords
			if j > 0 {
				choiceTwo = dp[i][j-1]
			}

			dp[i][j] = min(choiceOne, choiceTwo) + grid[i][j]

		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
