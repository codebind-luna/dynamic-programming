package main

var mod = 1000000007

func findWays(arr []int, k int) int {
	// Using Tabulation : Bottom up Approach

	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	if arr[0] <= k {
		dp[0][arr[0]] = 1
	}

	dp[0][0] = 2
	if arr[0] != 0 {
		dp[0][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			notTake := dp[i-1][j]
			take := 0

			if arr[i] <= j {
				take = dp[i-1][j-arr[i]]
			}

			dp[i][j] = (take + notTake) % mod
		}
	}

	return dp[n-1][k]
}
