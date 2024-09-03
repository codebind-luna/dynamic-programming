package main

import (
	"container/heap"
	"math"
)

type pair struct {
	first, second int
}

type priorityQueue []pair

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].first < pq[j].first
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(pair))
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Create adjacency list
	adj := make([][]pair, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], pair{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], pair{e[0], e[2]})
	}

	// Initialize priority queue
	pq := &priorityQueue{}
	heap.Init(pq)

	var cityno, mincitycount = 0, math.MaxInt
	for i := 0; i < n; i++ {
		dist := make([]int, n)
		for j := range dist {
			dist[j] = math.MaxInt
		}
		dist[i] = 0
		heap.Push(pq, pair{0, i})

		for pq.Len() > 0 {
			item := heap.Pop(pq).(pair)
			distance, node := item.first, item.second
			if distance > dist[node] {
				continue
			}
			for _, it := range adj[node] {
				if distance+it.second < dist[it.first] {
					dist[it.first] = distance + it.second
					heap.Push(pq, pair{dist[it.first], it.first})
				}
			}
		}

		count := 0
		for j := range dist {
			if j != i && dist[j] <= distanceThreshold {
				count++
			}
		}
		if count <= mincitycount {
			mincitycount = count
			cityno = i
		}
	}
	return cityno
}
