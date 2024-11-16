package main

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n+1)
	dp[n] = 0
	dp[n-1] = min(costs[0], min(costs[1], costs[2]))

	for i := n - 2; i >= 0; i-- {
		// buy 1 day pass
		option1 := costs[0] + dp[i+1]
		// buy 7 day pass
		option2 := costs[1] + dp[nextIdx(days, days[i]+7-1)]
		// buy 30 day pass
		option3 := costs[2] + dp[nextIdx(days, days[i]+30-1)]

		dp[i] = min(option1, min(option2, option3))
	}

	return dp[0]
}

func nextIdx(days []int, day int) int {
	s, e := 0, len(days)-1
	ans := len(days)
	for s <= e {
		mid := s + ((e - s) / 2)

		if days[mid] > day {
			ans = mid
			e = mid - 1
		} else {
			s = mid + 1
		}
	}

	return ans
}
