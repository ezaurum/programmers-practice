package main

import "testing"

func TestSolution(t *testing.T) {
	first := []int{
		70, 60, 55, 43, 57, 60, 44, 50,
	}
	second := []int{
		58, 40, 47, 90, 45, 52, 80, 40,
	}
	r := solution(8, 100, first, second)
	if r != 11 {
		t.Errorf("Expect 11, but got %d", r)
	}
}
