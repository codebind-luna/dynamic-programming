package main

import "math"

func minSubsetSumDifference(arr []int, n int) int {
	k := 0
	for i := 0; i < n; i++ {
		k += arr[i]
	}
	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, k+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	dp[0][arr[0]] = true

	for i := 1; i < n; i++ {
		for j := 1; j < k+1; j++ {
			if arr[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	minDiff := math.MaxInt

	for i := 0; i <= k/2; i++ {
		if dp[n-1][i] {
			minDiff = min(minDiff, abs((k-i)-i))
		}
	}
	return minDiff
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
