package main

import "math"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		dp[n-1][j] = matrix[n-1][j]
	}

	for i := n - 2; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			ld := math.MaxInt
			if j > 0 {
				ld = matrix[i][j] + dp[i+1][j-1]
			}

			rd := math.MaxInt

			if j < n-1 {
				rd = matrix[i][j] + dp[i+1][j+1]
			}

			d := matrix[i][j] + dp[i+1][j]

			dp[i][j] = min(min(ld, rd), d)
		}
	}

	res := math.MaxInt
	for j := 0; j < m; j++ {
		res = min(res, dp[0][j])
	}

	return res
}
