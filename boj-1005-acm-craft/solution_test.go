package main

import "testing"

func TestSolution(t *testing.T) {
	r := solution(4, 4, [][]int{{1, 2}, {1, 3}, {3, 4}, {2, 4}}, 4, []int{10, 1, 100, 10})
	if r != 120 {
		t.Errorf("Expect 120, but got %d", r)
	}
}

func TestSolution2(t *testing.T) {
	r := solution(8, 8, [][]int{{1, 2}, {1, 3}, {2, 5}, {2, 4}, {3, 6}, {5, 7}, {6, 7}, {7, 8}}, 7, []int{10, 20, 1, 5, 8, 7, 1, 43})
	if r != 39 {
		t.Errorf("Expect 120, but got %d", r)
	}
}

func TestSolution3(t *testing.T) {
	n := 5
	k := 10
	rules := [][]int{{4, 5}, {3, 5}, {3, 4}, {2, 5}, {2, 4}, {2, 3}, {1, 5}, {1, 4}, {1, 3}, {1, 2}}
	w := 4
	times := []int{100000, 99999, 99997, 99994, 99990}
	r := solution(n, k, rules, w, times)
	if r != 399990 {
		t.Errorf("Expect 399990, but got %d", r)
	}
}
