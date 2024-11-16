package main

func countPartitions(n, d int, arr []int) int {
	// Write your code here.
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	target := ((sum - d) / 2)
	mod := 1000_000_007
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < target+1; j++ {
			if i == 0 {
				if j == 0 && arr[0] == 0 {
					dp[0][j] = 2
				} else if j == 0 || arr[0] == j {
					dp[0][j] = 1
				}
			} else {
				notPick := dp[i-1][j]

				pick := 0

				if arr[i] <= j {
					pick = dp[i-1][j-arr[i]]
				}

				dp[i][j] = (notPick + pick) % mod
			}
		}
	}

	return dp[n-1][target]
}
