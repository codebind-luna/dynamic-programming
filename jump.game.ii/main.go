package main

import "math"

func minJump(i int, nums []int) int {
	if i == len(nums)-1 {
		return 0
	}

	if nums[i] == 0 {
		return 1e9
	}

	jump := nums[i]
	mJump := math.MaxInt
	for j := 1; j <= jump; j++ {
		if i+j <= len(nums)-1 {
			mJump = min(mJump, 1+minJump(i+j, nums))
		}
	}
	return mJump
}

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		if nums[i] == 0 {
			dp[i] = 1e9
			continue
		}

		minJ := math.MaxInt
		for j := 1; j <= nums[i]; j++ {
			if j+i <= n-1 {
				minJ = min(minJ, 1+dp[i+j])
			}
		}

		dp[i] = minJ
	}

	return dp[0]
}
