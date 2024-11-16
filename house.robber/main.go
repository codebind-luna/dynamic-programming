package main

func rob(nums []int) int {
	n := len(nums)

	dp := make([]int, n+1)
	dp[n] = 0
	dp[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}
	return dp[0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
