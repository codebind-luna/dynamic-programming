package main

func perfectSum(arr []int, n, sum int) int {
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, sum+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 1; j < sum+1; j++ {
			if i == 0 {
				if arr[i] == j {
					dp[i][j] = 1
				}
				continue
			}

			nonPick := dp[i-1][j]

			pick := 0

			if arr[i] <= j {
				pick = dp[i-1][j-arr[i]]
			}

			dp[i][j] = nonPick + pick
		}
	}

	return dp[n-1][sum]
}
