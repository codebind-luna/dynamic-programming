package main

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := [][]int{}

	n := len(startTime)

	for i := 0; i < n; i++ {
		jobs = append(jobs, []int{startTime[i], endTime[i], profit[i]})
	}

	// sort the jobs based on start time
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	dp := make([]int, n+1)
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		nonPick := dp[i+1]
		pick := jobs[i][2] + dp[nextIdx(jobs, i)]
		dp[i] = max(nonPick, pick)
	}
	return dp[0]
}

func nextIdx(jobs [][]int, currIdx int) int {
	s, e := 0, len(jobs)-1
	ans := len(jobs)

	for s <= e {
		mid := s + ((e - s) / 2)

		if jobs[mid][0] >= jobs[currIdx][1] {
			ans = mid
			e = mid - 1
		} else {
			s = mid + 1
		}
	}

	return ans
}
