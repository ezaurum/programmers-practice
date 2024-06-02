package main

// Solution https://www.acmicpc.net/problem/1697
func Solution(n, k int) int {
	type State struct {
		pos, time int
	}
	visited := make(map[int]bool)
	queue := []State{{n, 0}}
	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]
		if currentState.pos == k {
			return currentState.time
		}
		if visited[currentState.pos] {
			continue
		}
		visited[currentState.pos] = true
		if currentState.pos-1 >= 0 {
			queue = append(queue, State{currentState.pos - 1, currentState.time + 1})
		}
		if currentState.pos+1 <= 100000 {
			queue = append(queue, State{currentState.pos + 1, currentState.time + 1})
		}
		if currentState.pos*2 <= 100000 {
			queue = append(queue, State{currentState.pos * 2, currentState.time + 1})
		}

	}
	return -1
}
