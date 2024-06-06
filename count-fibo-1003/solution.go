package main

func solution(n int) (int, int) {
	if n == 0 {
		return 1, 0
	}
	if n == 1 {
		return 0, 1
	}

	zero, one := 0, 1
	for i := 2; i <= n; i++ {
		zero, one = one, zero+one
	}
	return zero, one
}
