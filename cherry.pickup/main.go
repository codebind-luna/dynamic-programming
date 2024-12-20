package main

import "math"

func cherryPickup(grid [][]int) int {
	n := len(grid)
	return max(0, cherryPickup2(grid, n, 0, 0, 0, 0))
}

func cherryPickup2(grid [][]int, n, r1, c1, r2, c2 int) int {
	// since we're only going down and to the right, no need to check for < 0
	// if we went out of the grid or hit a thorn, discourage this path by returning math.MinInt64
	if r1 >= n || c1 >= n || r2 >= n || c2 >= n || grid[r1][c1] == -1 || grid[r2][c2] == -1 {
		return math.MinInt64
	}

	// if person 1 reached the bottom right, return what's in there (could be 1 or 0)
	if r1 == n-1 && c1 == n-1 {
		return grid[r1][c1]
	}

	// if person 2 reached the bottom right, return what's in there (could be 1 or 0)
	if r2 == n-1 && c2 == n-1 {
		return grid[r2][c2]
	}

	var cherries int
	// if both persons standing on the same cell, don't double count and return what's in this cell (could be 1 or 0)
	if r1 == r2 && c1 == c2 {
		cherries = grid[r1][c1]
	} else {
		// otherwise, number of cherries collected by both of them equals the sum of what's on their cells
		cherries = grid[r1][c1] + grid[r2][c2]
	}

	// since each person of the 2 person can move only to the bottom or to the right, then the total number of cherries
	// equals the max of the following possibilities:
	//    P1     |      P2
	//   DOWN    |     DOWN
	//   DOWN    |     RIGHT
	//   RIGHT   |     DOWN
	//   RIGHT   |     RIGHT
	cherries += max(
		max(cherryPickup2(grid, n, r1+1, c1, r2+1, c2), cherryPickup2(grid, n, r1+1, c1, r2, c2+1)),
		max(cherryPickup2(grid, n, r1, c1+1, r2+1, c2), cherryPickup2(grid, n, r1, c1+1, r2, c2+1)),
	)

	return cherries
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
