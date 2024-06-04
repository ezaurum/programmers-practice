package main

import "testing"

func TestSolution4(t *testing.T) {
	n := 4
	r := solution(n)
	exp := 2
	if r != exp {
		t.Errorf("Expected %v, got %v\n", exp, r)
	}
}

func TestSolution5(t *testing.T) {
	n := 5
	r := solution(n)
	exp := 10
	if r != exp {
		t.Errorf("Expected %v, got %v\n", exp, r)
	}
}

func TestSolution15(t *testing.T) {
	n := 15
	r := solution(n)
	exp := 2279184
	if r != exp {
		t.Errorf("Expected %v, got %v\n", exp, r)
	}
}

func TestSolution2_15(t *testing.T) {
	n := 15
	r := solution2(n)
	exp := 2279184
	if r != exp {
		t.Errorf("Expected %v, got %v\n", exp, r)
	}
}
