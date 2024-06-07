package main

import "math"

// solution https://www.acmicpc.net/problem/1002
// 두 원의 교점의 개수를 구하는 문제
// 두 원의 중심 사이의 거리가 두 원의 반지름의 합과 같으면 두 원은 외접한다.
// 두 원의 중심 사이의 거리가 두 원의 반지름의 차이와 같으면 두 원은 내접한다.
// 두 원의 중심 사이의 거리가 두 원의 반지름의 합보다 크면 두 원은 만나지 않는다.
// 두 원의 중심 사이의 거리가 두 원의 반지름의 차이보다 작으면 두 원은 만나지 않는다.
// 두 원의 중심 사이의 거리가 0이고 두 원의 반지름이 같으면 무한대로 만난다 = 동심원.
// 그 외의 경우는 두 원이 두 점에서 만난다.
func solution(n int, arr [][]int) []int {
	result := make([]int, n)
	for n > 0 {
		n--
		x1, y1, r1, x2, y2, r2 := arr[n][0], arr[n][1], arr[n][2], arr[n][3], arr[n][4], arr[n][5]
		distance := math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
		r := float64(r1 + r2)
		rs := max(float64(r1-r2), float64(r2-r1))
		if distance == 0 && r1 == r2 {
			result[n] = -1
		} else if distance == r || distance == rs {
			result[n] = 1
		} else if distance > r || rs > distance {
			result[n] = 0
		} else {
			result[n] = 2
		}
	}
	return result
}
