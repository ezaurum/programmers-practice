package main

import "testing"

func TestGraph(t *testing.T) {
	edges := [][]int{
		{2, 3},
		{4, 3},
		{1, 1},
		{2, 1},
	}
	result := solution(edges)
	if result[0] != 2 || result[1] != 1 || result[2] != 1 || result[3] != 0 {
		t.Errorf("Expected [2, 1, 1, 0], got %v", result)
	}
}
