package main

import (
	"math"
)

func solution(n int, vectors [][]int) float64 {
	minR := math.MaxFloat64
	type state struct {
		visited []bool
		vectors [][]int
	}

	stack := []state{{visited: make([]bool, n), vectors: make([][]int, 0)}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 다 채웠을 때
		if len(current.vectors) == n {
			cv := current.vectors
			i := 0
			var vv [][]int
			for i < n {
				x0 := cv[i][0]
				x1 := cv[i+1][0]
				vv = append(vv, []int{x0 - x1, cv[i][1] - cv[i+1][1]})
				i += 2
			}
			minR = math.Min(minR, checkLength(vv))
			continue
		}

		// next
		for i := 0; i < n; i++ {
			visited := current.visited
			if visited[i] {
				continue
			}
			newVisited := make([]bool, n)
			copy(newVisited, visited)
			newVector := make([][]int, len(current.vectors))
			copy(newVector, current.vectors)
			ints := vectors[i]
			newVisited[i] = true
			newState := state{visited: newVisited, vectors: append(newVector, ints)}
			stack = append(stack, newState)
		}
	}

	return minR
}

func checkLength(vectors [][]int) float64 {
	sumX := 0
	sumY := 0
	for _, vector := range vectors {
		sumX += vector[0]
		sumY += vector[1]
	}
	return math.Sqrt(float64(sumX*sumX + sumY*sumY))
}
