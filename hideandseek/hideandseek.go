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

const MAX = 100001

// Solution2 이래저래 애를 써 봤는데 실제로 빨라지지 않았다 - 최적화를 go가 해주는 것 같다
// https://www.acmicpc.net/problem/1697
func Solution2(n, k int) int {
	visited := make(map[int]bool, MAX)
	queue := [][]int{{n, 0}}
	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]
		curPosition := currentState[0]
		curTime := currentState[1]
		visited[curPosition] = true
		if curPosition == k {
			return curTime
		}
		next := []int{curPosition - 1, curPosition + 1, curPosition * 2}
		for _, nn := range next {
			if nn >= 0 && nn < MAX && !visited[nn] {
				queue = append(queue, []int{nn, curTime + 1})
			}
		}
	}
	return -1
}
