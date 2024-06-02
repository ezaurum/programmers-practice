package main

import "testing"

func TestSolution(t *testing.T) {
	n := 4
	m := 5
	start := 1
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 4},
		{3, 4},
	}
	r := Solution(n, m, start, edges)
	if r[0][0] != "1" || r[0][1] != "2" || r[0][2] != "4" || r[0][3] != "3" {
		t.Error("Expected [1 2 4 3], got ", r[0])
	}
	if r[1][0] != "1" || r[1][1] != "2" || r[1][2] != "3" || r[1][3] != "4" {
		t.Error("Expected [1 2 3 4], got ", r[1])
	}
}

func TestSolution2(t *testing.T) {

	n := 5
	m := 5
	start := 3
	edges := [][]int{
		{5, 4},
		{5, 2},
		{1, 2},
		{3, 4},
		{3, 1},
	}
	r := Solution(n, m, start, edges)
	if r[0][0] != "3" || r[0][1] != "1" || r[0][2] != "2" || r[0][3] != "5" || r[0][4] != "4" {
		t.Error("Expected [3 1 2 5 4], got ", r[0])
	}
	if r[1][0] != "3" || r[1][1] != "1" || r[1][2] != "4" || r[1][3] != "2" || r[1][4] != "5" {
		t.Error("Expected [3 1 4 2 5], got ", r[1])
	}
}
