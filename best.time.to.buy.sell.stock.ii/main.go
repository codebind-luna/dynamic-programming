package main

func f(idx int, buy int, prices []int, dp [][]int) int {
	if idx == len(prices) {
		return 0
	}
	if dp[idx][buy] != -1 {
		return dp[idx][buy]
	}
	profit := 0
	if buy == 1 {
		profit = max(-prices[idx]+f(idx+1, 0, prices, dp), 0+f(idx+1, 1, prices, dp))
	} else {
		profit = max(prices[idx]+f(idx+1, 1, prices, dp), 0+f(idx+1, 0, prices, dp))
	}

	dp[idx][buy] = profit
	return dp[idx][buy]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// memoization
// func maxProfit(prices []int) int {
// 	dp := make([][]int, len(prices))
// 	for i := 0; i < len(prices); i++ {
// 		dp[i] = make([]int, 2)
// 	}

// 	for i := 0; i < len(prices); i++ {
// 		for j := 0; j < 2; j++ {
// 			dp[i][j] = -1
// 		}
// 	}
// 	return f(0, 1, prices, dp)
// }

// tabulation
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 2)
	}

	dp[n][0], dp[n][1] = 0, 0

	for idx := n - 1; idx >= 0; idx-- {
		for buy := 0; buy < 2; buy++ {
			profit := 0
			if buy == 1 {
				profit = max(-prices[idx]+dp[idx+1][0], 0+dp[idx+1][1])
			} else {
				profit = max(prices[idx]+dp[idx+1][1], 0+dp[idx+1][0])
			}

			dp[idx][buy] = profit
		}
	}

	return dp[0][1]
}
