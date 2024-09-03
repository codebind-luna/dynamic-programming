package main

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
		for j := 0; j < k+1; j++ {
			if j == 0 {
				dp[i][j] = true
				continue
			}

			if i == 0 {
				dp[i][j] = (nums[i] == j)
				continue
			}

			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n-1][k]
}
