package main

import "math"

func mPoints(i, j int, points [][]int, dp [][]int) int {
	if i == len(points) {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	maxP := math.MinInt
	for k := 0; k < len(points[0]); k++ {
		p := points[i][k]

		p -= abs(j - k)
		maxP = max(maxP, p+mPoints(i+1, k, points, dp))

	}
	dp[i][j] = maxP
	return dp[i][j]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxPoints(points [][]int) int64 {
	dp := make([][]int, len(points))

	for i := 0; i < len(points); i++ {
		dp[i] = make([]int, len(points[0]))
	}

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			dp[i][j] = -1
		}
	}
	return int64(mPoints(0, 0, points, dp))
}
