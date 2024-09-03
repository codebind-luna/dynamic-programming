package main

import "math"

func knapsack(weight []int, value []int, n, maxWeight int) int {
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, maxWeight+1)
	}

	for j := 0; j < maxWeight+1; j++ {
		if weight[0] <= j {
			dp[0][j] = value[0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < maxWeight+1; j++ {
			nonPick := dp[i-1][j]
			pick := math.MinInt

			if weight[i] <= j {
				pick = value[i] + dp[i-1][j-weight[i]]
			}

			dp[i][j] = max(nonPick, pick)
		}
	}

	return dp[n-1][maxWeight]
}
