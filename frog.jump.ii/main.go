package main

import "math"

func frogJump(n int, k int, height []int) int {
	dp := make([]int, n)

	dp[0] = 0

	for i := 1; i < n; i++ {
		minStep := math.MaxInt

		for p := 1; p <= k; k++ {
			if i-p >= 0 {
				ans := dp[i-p] + abs(height[i]-height[i-p])
				minStep = min(minStep, ans)
			}
		}
		dp[i] = minStep
	}

	return dp[n-1]
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
