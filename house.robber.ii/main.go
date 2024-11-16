package main

func helper(nums []int) int {
	n := len(nums)

	dp := make([]int, n+1)
	dp[n] = 0
	dp[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}
	return dp[0]
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	numsWithoutFirst := nums[1:]
	numsWithoutLast := nums[:len(nums)-1]

	return max(helper(numsWithoutFirst), helper(numsWithoutLast))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
