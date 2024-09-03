package main

import "math"

func frogJump(n int, heights []int) int {
	dp := make([]int, n)

	dp[0] = 0

	for i := 1; i < n; i++ {
		s1 := dp[i-1] + abs(heights[i]-heights[i-1])

		s2 := math.MaxInt
		if i > 1 {
			s2 = dp[i-2] + abs(heights[i]-heights[i-2])
		}

		dp[i] = min(s1, s2)
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
