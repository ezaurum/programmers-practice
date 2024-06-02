package main

import "testing"

func TestSolution(t *testing.T) {
	n := 5
	k := 17
	r := Solution(n, k)
	if r != 4 {
		t.Error("Expected 4, got ", r)
	}
}
