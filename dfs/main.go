package main

// package main

// import "log"

// func dfs(node int, visited map[int]bool, adjacencyList map[int][]int) {
// 	if visited[node] {
// 		return
// 	}

// 	log.Println(node)

// 	visited[node] = true

// 	for _, v := range adjacencyList[node] {
// 		if !visited[v] {
// 			dfs(v, visited, adjacencyList)
// 		}
// 	}

// }

// func main() {
// 	adjacencyList := map[int][]int{
// 		1: {2, 4},
// 		2: {1, 4},
// 		4: {1, 2},
// 	}
// 	visited := make(map[int]bool)
// 	nodes := []int{1, 2, 4}

// 	for i := 0; i < len(nodes); i++ {
// 		if !visited[nodes[i]] {
// 			dfs(nodes[i], visited, adjacencyList)
// 		}
// 	}

// 	log.Println(visited)
// }
