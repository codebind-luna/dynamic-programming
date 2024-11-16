package main

import "math"

func minCoins(idx, amount int, coins []int) int {
	if idx == 0 {
		if amount%coins[0] == 0 {
			return amount / coins[0]
		}
		return math.MaxInt
	}

	if amount == 0 {
		return 0
	}

	notTake := minCoins(idx-1, amount, coins)
	take := math.MaxInt

	if coins[idx] <= amount {
		take = 1 + minCoins(idx, amount-coins[idx], coins)
	}
	return min(notTake, take)
}

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, amount+1)
	}

	for j := 0; j < amount+1; j++ {
		if j%coins[0] == 0 {
			dp[0][j] = j / coins[0]
		} else {
			dp[0][j] = 1e9
		}
	}

	for idx := 1; idx < n; idx++ {
		for j := 1; j <= amount; j++ {
			notTake := dp[idx-1][j]
			take := math.MaxInt

			if coins[idx] <= j {
				take = 1 + dp[idx][j-coins[idx]]
			}
			dp[idx][j] = min(notTake, take)
		}
	}

	if dp[n-1][amount] >= 1e9 {
		return -1
	}
	return dp[n-1][amount]
}
