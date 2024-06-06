package main

import "testing"

func TestSolution(t *testing.T) {
	n := 3
	z, o := solution(n)
	if z != 1 || o != 2 {
		t.Errorf("Expected 1, 2, got %v, %v", z, o)
	}
}
