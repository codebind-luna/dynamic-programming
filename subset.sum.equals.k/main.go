package main

func subsetSumToK(n int, k int, arr []int) bool {
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, k+1)
	}

	// empty subset i.e {} will result 0 sum
	// hence no matter the nums of elements in the positive
	// array 0 subset sum can be found
	for i := 1; i < n; i++ {
		dp[i][0] = true
	}

	for j := 1; j < k+1; j++ {
		dp[0][j] = (arr[0] == j)
	}

	dp[0][0] = true

	for i := 1; i < n; i++ {
		for j := 1; j < k+1; j++ {
			nonPick := dp[i-1][j]

			pick := false
			if arr[i] <= j {
				pick = dp[i-1][j-arr[i]]
			}

			dp[i][j] = nonPick || pick
		}
	}

	return dp[n-1][k]
}
