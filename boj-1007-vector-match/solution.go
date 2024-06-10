package main

import (
	"math"
)

func solution(n int, vectors [][]int) float64 {
	// 합산해 놓고, 절반씩 나눠서 계산
	totalSumX := 0
	totalSumY := 0
	for _, vector := range vectors {
		totalSumX += vector[0]
		totalSumY += vector[1]
	}
	minR := math.MaxFloat64
	type state struct {
		index    int
		sumX     int
		sumY     int
		selected int
	}

	stack := []state{{}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 반 채웠을 때 - 나머지는 반대편에 넣어야 함
		if current.selected == n/2 {
			tx := totalSumX - current.sumX*2
			ty := totalSumY - current.sumY*2
			minR = math.Min(minR, math.Sqrt(float64(tx*tx+ty*ty)))
			continue
		}
		if current.index == n {
			continue
		}
		ints := vectors[current.index]

		stack = append(stack,
			state{index: current.index + 1, sumX: current.sumX + ints[0], sumY: current.sumY + ints[1], selected: current.selected + 1})
		stack = append(stack,
			state{index: current.index + 1, sumX: +current.sumX, sumY: current.sumY, selected: current.selected})
	}

	return minR
}
