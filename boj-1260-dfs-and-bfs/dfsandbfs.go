package main

import "fmt"
import "sort"

// Solution https://www.acmicpc.net/problem/1260
func Solution(n int, m int, start int, edges [][]int) [][]string {
	graph := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	// 정방향으로 정렬해 둔다
	for k, v := range graph {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		graph[k] = v
	}

	dfsResult := dfs(start, graph)
	bfsResult := bfs(start, graph)

	return [][]string{dfsResult, bfsResult}
}

func bfs(start int, graph map[int][]int) []string {
	queue := []int{start}
	visited := make(map[int]bool)
	var result []string
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if visited[n] {
			continue
		}
		visited[n] = true
		for _, neighbor := range graph[n] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
		result = append(result, fmt.Sprint(n))
	}
	return result
}

func dfs(start int, graph map[int][]int) []string {
	stack := []int{start}
	visited := make(map[int]bool)
	var result []string
	for len(stack) > 0 {
		// pop 구현
		last := len(stack) - 1
		n := stack[last]
		stack = stack[:last]
		visited[n] = true
		result = append(result, fmt.Sprint(n))

		neighbors := append([]int{}, graph[n]...)
		for len(neighbors) > 0 {
			// pop
			lastN := len(neighbors) -1
			nn := neighbors[lastN]
			neighbors = neighbors[:lastN]
			if visited[nn] {
				continue
			}
			stack = append(stack, nn)
		}
	}
	return result
}

