package main

func f(idx int, buy int, cap int, prices []int, dp [][][]int) int {
	if cap == 0 {
		return 0
	}

	if idx == len(prices) {
		return 0
	}

	if dp[idx][buy][cap] != -1 {
		return dp[idx][buy][cap]
	}

	profit := 0
	if buy == 1 {
		profit = max(-prices[idx]+f(idx+1, 0, cap, prices, dp), 0+f(idx+1, 1, cap, prices, dp))
	} else {
		profit = max(prices[idx]+f(idx+1, 1, cap-1, prices, dp), 0+f(idx+1, 0, cap, prices, dp))
	}

	dp[idx][buy][cap] = profit
	return dp[idx][buy][cap]
}

// Memoization
func maxProfitMemo(k int, prices []int) int {
	n := len(prices)
	dp := make([][][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}

	for i := 0; i < len(prices); i++ {
		for j := 0; j < 2; j++ {
			for l := 0; l <= k; l++ {
				dp[i][j][l] = -1
			}
		}
	}
	return f(0, 1, k, prices, dp)
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp := make([][][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}

	for j := 0; j < 2; j++ {
		for k1 := 0; k1 <= k; k1++ {
			dp[n][j][k1] = 0
		}
	}

	for i := 0; i <= n; i++ {
		for j := 0; j < 2; j++ {
			dp[i][j][0] = 0
		}
	}

	for idx := n - 1; idx >= 0; idx-- {
		for buy := 0; buy < 2; buy++ {
			for cap := 1; cap <= k; cap++ {
				profit := 0
				if buy == 1 {
					profit = max(-prices[idx]+dp[idx+1][0][cap], 0+dp[idx+1][1][cap])
				} else {
					profit = max(prices[idx]+dp[idx+1][1][cap-1], 0+dp[idx+1][0][cap])
				}
				dp[idx][buy][cap] = profit
			}
		}
	}

	return dp[0][1][k]
}
