package main

func MinInsertions(s string) int {
	n := len(s)

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if i > j {
				dp[i][j] = 0
				continue
			}

			if i == j {
				dp[i][j] = 0
				continue
			}

			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
				continue
			}

			dp[i][j] = min(1+dp[i+1][j], 1+dp[i][j-1])
		}
	}

	return dp[0][n-1]
}
