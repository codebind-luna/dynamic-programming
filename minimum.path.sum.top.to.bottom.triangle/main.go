package main

import (
	"fmt"
	"math"
) // to get max int value

func minimumTotal(triangle [][]int) int {
	/*
	   Create dp array:
	       dp[i][j] contains the minimum path sum to get from
	       triangle[0][0] to triangle[i][j] inclusively.
	*/
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		// i'th row has i+1 elements in it
		dp[i] = make([]int, i+1)
	}

	// As the 0th row has no preceding rows, calculate it by hand.
	dp[0][0] = triangle[0][0]

	// minPathSum would be the final answer
	minPathSum := math.MaxInt32

	// Calculate the whole triangle in this loop
	for i := 1; i < m; i++ {
		for j := 0; j < len(dp[i]); j++ {
			/*
			   The avoid accessing slices at indices -1 and len(dp[i-1]),
			   we have to look at two corner cases:
			   1. We can get to 0'th element of the row only from the 0'th element of the previous row.
			   2. We can get to the last element of the rown only from the last element of the prev. row.
			   3. Otherwise, current minimum path to the triangle[i][j] equals to
			       the min path between the two elements of the row above. We are looking at only two elements,
			       because we can't get to the current element from ANY element of the prev. row.

			       So, dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			*/
			switch j {
			case 0:
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			case len(dp[i]) - 1:
				dp[i][j] = dp[i-1][len(dp[i-1])-1] + triangle[i][j]
			default:
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	// for i := 0; i < m; i++ {
	// 	for j := 0; j < i+1; j++ {
	// 		fmt.Printf("%d ", dp[i][j])
	// 	}
	// 	fmt.Println()
	// }

	// The last dp row contains all the paths. Now we just need to find the minimum path among them.
	for _, lastRowSum := range dp[len(dp)-1] {
		minPathSum = min(lastRowSum, minPathSum)
	}

	return minPathSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}}
	fmt.Println(minimumTotal(triangle))
}
