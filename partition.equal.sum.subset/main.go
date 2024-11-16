package main

// input [1,2,5] will output false
// input [2,3,5] will output true
func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	k := sum / 2

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, k+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	for j := 0; j < k+1; j++ {
		dp[0][j] = (nums[0] == j)
	}

	for i := 1; i < n; i++ {
		for j := 1; j < k+1; j++ {
			nonPick := dp[i-1][j]
			pick := false
			if nums[i] <= j {
				pick = dp[i-1][j-nums[i]]
			}
			dp[i][j] = nonPick || pick
		}
	}

	return dp[n-1][k]
}
