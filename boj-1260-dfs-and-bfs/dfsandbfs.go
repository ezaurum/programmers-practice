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

	dfsResult := dfs(start, graph)
	bfsResult := bfs(start, graph)

	return [][]string{dfsResult, bfsResult}
}

func bfs(start int, graph map[int][]int) []string {
	for k, v := range graph {
		// bfs는 정방향 정렬
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		graph[k] = v
	}
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
	for k, v := range graph {
		// dfs는 역순으로 정렬
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
		graph[k] = v
	}
	stack := []int{start}
	visited := make(map[int]bool)
	var result []string
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[n] {
			continue
		}
		visited[n] = true
		for _, neighbor := range graph[n] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
		result = append(result, fmt.Sprint(n))
	}
	return result
}
