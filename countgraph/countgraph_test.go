package main

import (
	"testing"
)

func dfs(node int, graph map[int][]int, visited map[int]bool) (int, int) {
	stack := []int{node}
	// Check if the graph is disconnected
	if len(graph[node]) == 0 {
		return -1, -1
	}
	nodes, edges := 0, 0
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[n] {
			nodes += 1
			edges += len(graph[n])
			visited[n] = true
			for _, neighbor := range graph[n] {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
	return nodes, edges
}

// Main solution function
func solution(edges [][]int) []int {
	graph := make(map[int][]int)
	indegree := make(map[int]int)
	outdegree := make(map[int]int)
	nodes := make(map[int]bool)

	// Build the graph
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		outdegree[a]++
		indegree[b]++
		nodes[a] = true
		nodes[b] = true
	}

	donutCount, barCount, eightCount := 0, 0, 0
	startNode := -1
	visited := make(map[int]bool)

	// Identify the generated node
	for node := range nodes {
		if indegree[node] == 0 && outdegree[node] > 1 {
			startNode = node
			arr := graph[node]
			var narr []int
			for _, n := range arr {
				if outdegree[n] == 0 && indegree[n] == 1 {
					narr = append(narr, n)
				}
			}
			graph[node] = narr
			break
		}
	}

	// Explore the graph components

	for node := range nodes {
		if visited[node] {
			continue
		}
		n, e := dfs(node, graph, visited)
		if n == -1 {
			continue
		} else if n == e {
			donutCount++
		} else if n == e+1 {
			barCount++
		} else {
			eightCount++
		}
	}

	return []int{startNode, donutCount, barCount, eightCount}
}

func TestGraph(t *testing.T) {
	edges := [][]int{
		{2, 3},
		{4, 3},
		{1, 1},
		{2, 1},
	}
	result := solution(edges)
	if result[0] != 2 || result[1] != 1 || result[2] != 1 || result[3] != 0 {
		t.Errorf("Expected [2, 1, 1, 0], got %v", result)
	}
}
