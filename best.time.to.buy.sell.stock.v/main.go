package main

func f(idx int, buy int, prices []int, dp [][]int) int {
	if idx >= len(prices) {
		return 0
	}
	if dp[idx][buy] != -1 {
		return dp[idx][buy]
	}
	profit := 0
	if buy == 1 {
		profit = max(-prices[idx]+f(idx+1, 0, prices, dp), 0+f(idx+1, 1, prices, dp))
	} else {
		profit = max(prices[idx]+f(idx+2, 1, prices, dp), 0+f(idx+1, 0, prices, dp))
	}

	dp[idx][buy] = profit
	return dp[idx][buy]
}

func fRec(idx int, buy int, prices []int) int {
	if idx == len(prices) {
		return 0
	}

	profit := 0
	if buy == 1 {
		profit = max(-prices[idx]+fRec(idx+1, 0, prices), 0+fRec(idx+1, 1, prices))
	} else {
		profit = max(prices[idx]+fRec(idx+2, 1, prices), 0+fRec(idx+1, 0, prices))
	}

	return profit
}

// func maxProfit(prices []int) int {
// 	return fRec(0, 1, prices)
// }

// Memoization
// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	dp := make([][]int, n)
// 	for i := 0; i < len(prices); i++ {
// 		dp[i] = make([]int, 2)
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < 2; j++ {
// 			dp[i][j] = -1
// 		}
// 	}
// 	return f(0, 1, prices, dp)
// }

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, 2)
	}

	// dp[n][0], dp[n][1] = 0, 0

	for idx := n - 1; idx >= 0; idx-- {
		for buy := 0; buy < 2; buy++ {
			profit := 0
			if buy == 1 {
				profit = max(-prices[idx]+dp[idx+1][0], 0+dp[idx+1][1])
			} else {
				profit = max(prices[idx]+dp[idx+2][1], 0+dp[idx+1][0])
			}

			dp[idx][buy] = profit
		}
	}

	return dp[0][1]
}
