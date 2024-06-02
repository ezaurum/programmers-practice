package missle

import (
	"sort"
)

// solution https://school.programmers.co.kr/learn/courses/30/lessons/181188
func solution(targets [][]int) int {
	sort.Slice(targets, func(i, j int) bool {
		return targets[i][1] < targets[j][1]
	})

	currentEnd := -1
	count := 0
	for _, target := range targets {
		start := target[0]
		end := target[1]
		if currentEnd <= start {
			count++
			currentEnd = end
		}
	}
	return count
}
