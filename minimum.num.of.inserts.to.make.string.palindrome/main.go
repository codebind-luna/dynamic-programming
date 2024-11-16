package main

func minsteps(i, j int, s string, dp [][]int) int {
	if i > j {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	if s[i] == s[j] {
		dp[i][j] = minsteps(i+1, j-1, s, dp)
	} else {
		dp[i][j] = min(1+minsteps(i+1, j, s, dp), 1+minsteps(i, j-1, s, dp))
	}

	return dp[i][j]
}

func minInsertions(s string) int {
	n := len(s)

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return minsteps(0, n-1, s, dp)
}
