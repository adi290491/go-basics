package graph

import (
	"fmt"
)

func graphdemo() {

	list := [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 1}}
	var adjList = AdjacencyList(list)
	fmt.Println(adjList)

	fmt.Printf("******Depth First Search******\n")
	fmt.Println(dfs(adjList))
	fmt.Printf("******Breadth First Search******\n")
	fmt.Println(bfs(adjList))
}

func AdjacencyList(list [][]int) map[int][]int {
	adjList := map[int][]int{}
	for _, pair := range list {
		adjList[pair[0]] = append(adjList[pair[0]], pair[1])
		adjList[pair[1]] = append(adjList[pair[1]], pair[0])
	}

	return adjList
}

var graph = make(map[int]map[int]bool)

func addEdge(from, to int) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[int]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func dfs(adjList map[int][]int) []int {
	var seen = map[int]bool{}
	var stack = []int{}
	stack = append(stack, 0)
	answer := []int{}
	for len(stack) > 0 {
		var node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if seen[node] {
			continue
		}

		answer = append(answer, node)
		seen[node] = true
		for _, neighbor := range adjList[node] {
			stack = append(stack, neighbor)
		}
	}

	return answer
}

func bfs(adjList map[int][]int) []int {
	var seen = map[int]bool{}
	var queue = []int{}
	queue = append(queue, 0)
	answer := []int{}
	for len(queue) > 0 {
		var node = queue[0]
		queue = queue[1:]
		if seen[node] {
			continue
		}

		answer = append(answer, node)
		seen[node] = true
		for _, neighbor := range adjList[node] {
			queue = append(queue, neighbor)
		}
	}

	return answer
}
