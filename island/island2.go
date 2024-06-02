package missle

import (
	"sort"
)

func checkConnected(graph [][]int, visited []bool, from, to int) bool {
	// 이미 방문한 노드인 경우 = 사이클이 생기는 경우
	if visited[from] {
		return false
	}
	visited[from] = true
	for _, neighbor := range graph[from] {
		if visited[neighbor] {
			continue
		}
		if neighbor == to {
			return true
		}
		if checkConnected(graph, visited, neighbor, to) {
			return true
		}
	}
	return false
}

func Solution2(n int, costs [][]int) int {
	// 코스트로 정렬
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})
	graph := make([][]int, n)
	sum := 0
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		to := cost[1]
		from := cost[0]
		// from => to 로 가는 경로만 찾는 게 목적이므로 visited는 계속 초기화
		visited := make([]bool, n)
		if checkConnected(graph, visited, from, to) {
			continue
		}
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
		// 비용 추가
		sum += cost[2]
	}
	return sum
}
