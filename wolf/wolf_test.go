package wolf

import (
	"testing"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/92343#
func Solution(info []int, edges [][]int) int {
	graph := make([][]int, len(info))
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
	}

	maxSheep := 0
	type State struct {
		wolf  int
		sheep int
		// 현재까지 패스
		path    []int
		visited []bool
	}

	// 백트래킹을 위한 스택
	stack := []State{{0, 0,
		// 0번 노드부터 시작하는 경로
		[]int{0}, make([]bool, len(info))}}

	// 스택이 비어있지 않을 때까지 반복
	for len(stack) > 0 {

		// 현재 state pop, 마지막 요소를 가져오고 슬라이스에서 제거
		currentState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if currentState.sheep > maxSheep {
			maxSheep = currentState.sheep
		}

		// 현재 path
		currentPath := currentState.path

		// 트리는 방문 여부를 체크 할 필요가 없다. 한 번 만 갈 수 있기 때문
		// 하지만 이 문제의 경우는 다시 돌아가는 경우가 있기 때문에 방문 여부를 체크해야 한다.
		for i, next := range currentPath {
			// 이동할 수 있는 노드를 추가 = 이웃 노드와 자식 노드
			// 순서는 이웃 노드를 먼저 추가해야 한다. DFS이기 때문
			// 0. 자신을 제외한 이웃 노드를 추가 - 그래야 백트래킹이 가능
			//
			p := append([]int{}, currentPath[:i]...)
			p = append(p, currentPath[i+1:]...)
			// 1. 자식 노드를 추가
			p = append(p, graph[next]...)

			newState := State{
				wolf:  currentState.wolf,
				sheep: currentState.sheep,
				path:  p,
			}
			if info[next] == 0 {
				newState.sheep += 1
			} else {
				newState.wolf += 1
			}
			// 늑대가 양보다 많으면 스킵
			if newState.wolf >= newState.sheep {
				continue
			}
			// 다음 노드를 스택에 추가
			stack = append(stack, newState)
		}
	}

	return maxSheep
}

func TestSolution(t *testing.T) {
	info := []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1}
	edges := [][]int{{0, 1}, {1, 2}, {1, 4}, {0, 8}, {8, 7}, {9, 10}, {9, 11}, {4, 3}, {6, 5}, {4, 6}, {8, 9}}
	result := Solution(info, edges)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestSolution2(t *testing.T) {
	info := []int{0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0}
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}, {3, 7}, {4, 8}, {6, 9}, {9, 10}}
	result := Solution(info, edges)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestSolution3(t *testing.T) {
	info := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}, {3, 7}, {4, 8}, {6, 9}, {9, 10}}
	result := Solution(info, edges)
	v := 11
	if result != v {
		t.Errorf("Expected %v, got %v", v, result)
	}
}
