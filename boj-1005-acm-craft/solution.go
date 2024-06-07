package main

func solution(n int, k int, rules [][]int, w int, times []int) int {
	// n+1을 하는 이유는 계산의 편의를 위해 0 인덱스를 사용하지 않기 위함
	// 인접 리스트, 크기 비교를 해서 작은 수(key) -> 큰 수(value)로 되게 저장할 예정

	// 그래프를 만들어서 위상정렬을 해야 함
	// 위상정렬을 하면서 시간을 계산하면 됨
	// 위상정렬을 하면서 동시에 짓는 건지 확인해야 함
	// 동시에 짓는 건지 확인하려면 그래프를 만들어야 함
	// index 0은 사용하지 않음
	adjacencyList := make([][]int, n+1)
	indegree := make([]int, n+1)
	for _, rule := range rules {
		a, b := rule[0], rule[1]
		adjacencyList[a] = append(adjacencyList[a], b)
		indegree[b]++
	}

	if indegree[w] == 0 {
		return times[w-1]
	}

	buildTime := make([]int, n+1)

	var queue []int
	for i, ind := range indegree {
		if ind == 0 && i != 0 {
			queue = append(queue, i)
			buildTime[i] = times[i-1]
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		neighbor := adjacencyList[node]

		for _, next := range neighbor {
			nextTime := times[next-1]
			if buildTime[next] < buildTime[node]+nextTime {
				buildTime[next] = buildTime[node] + nextTime
			}
			indegree[next]--
			if indegree[next] == 0 {
				if next == w {
					return buildTime[next]
				}
				queue = append(queue, next)
			}
		}
	}

	return buildTime[w]
}
