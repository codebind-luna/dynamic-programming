package main

func ninjaTraining(n int, points [][]int) int {
	dp := make([][]int, n)

	w := len(points[0])

	for i := 0; i < n; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < w+1; j++ {
			if i == 0 {
				maxP := 0
				for p := 0; p < w; p++ {
					if p != j {
						maxP = max(maxP, points[i][p])
					}
				}
				dp[i][j] = maxP
				continue
			}

			maxP := 0
			for p := 0; p < w; p++ {
				if p != j {
					ans := dp[i-1][p] + points[i][p]
					maxP = max(maxP, ans)
				}
			}

			dp[i][j] = maxP
		}
	}

	return dp[n-1][w]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
