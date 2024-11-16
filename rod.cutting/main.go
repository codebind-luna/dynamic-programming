package main

import "math"

func cutRod(price []int, n int) int {
	// Write your code here.
	m := len(price)

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = price[0] * n
	}

	for i := 1; i < m; i++ {
		for j := 1; j <= n; j++ {
			notTake := dp[i-1][j]
			take := math.MinInt
			if i+1 <= j {
				take = price[i] + dp[i][j-i+1]
			}
			dp[i][j] = max(take, notTake)
		}
	}

	return dp[m-1][n]
}
