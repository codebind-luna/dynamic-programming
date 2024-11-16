package main

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[n][0], dp[n][1] = 0, 0

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= 1; j++ {
			profit := 0
			if j == 1 {
				profit = max(-prices[i]+dp[i+1][0], dp[i+1][1])
			} else {
				profit = max((prices[i]-fee)+dp[i+1][1], dp[i+1][0])
			}

			dp[i][j] = profit
		}
	}

	return dp[0][1]
}
