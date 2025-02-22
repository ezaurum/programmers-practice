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
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[n] {
			continue
		}
		visited[n] = true
		neighbors := graph[n]
		// 스택에 넣을 때는 처음 방문할 곳을 마지막에 넣어야 하므로 거꾸로 꺼낸다.
		// => 이것도 pop하면 되겠다.
		// => 이래서 재귀구나
		for i := len(neighbors) - 1; i >= 0; i-- {
			neighbor := neighbors[i]
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
		result = append(result, fmt.Sprint(n))
	}
	return result
}
