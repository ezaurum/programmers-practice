package main

// solution https://www.acmicpc.net/problem/1874
func solution(targets []int) []string {
	var result []string
	var stack []int
	current := 1
	for _, targetNumber := range targets {
		for current <= targetNumber {
			stack = append(stack, current)
			result = append(result, "+")
			current++
		}
		if stack[len(stack)-1] == targetNumber {
			stack = stack[:len(stack)-1]
			result = append(result, "-")
		} else {
			return []string{"NO"}
		}
	}

	return result
}
