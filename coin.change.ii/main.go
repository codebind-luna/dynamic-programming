package main

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, amount+1)
	}

	for j := 1; j <= amount; j++ {
		if j%coins[0] == 0 {
			dp[0][j] = 1
		}
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	dp[0][0] = 1

	for idx := 1; idx < n; idx++ {
		for j := 1; j <= amount; j++ {
			notTake := dp[idx-1][j]
			take := 0

			if coins[idx] <= j {
				take = dp[idx][j-coins[idx]]
			}
			dp[idx][j] = take + notTake
		}
	}

	return dp[n-1][amount]
}
