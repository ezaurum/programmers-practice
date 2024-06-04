package main

import "testing"

func TestSolution(t *testing.T) {
	bridgeLength := 2
	weight := 10
	truckWeights := []int{7, 4, 5, 6}
	r := solution(bridgeLength, weight, truckWeights)
	exp := 8
	if r != exp {
		t.Errorf("Expected %v, got %v\n", exp, r)
	}
}

func TestSolution1(t *testing.T) {
	bridgeLength := 100
	weight := 100
	truckWeights := []int{10}
	r := solution(bridgeLength, weight, truckWeights)
	exp := 101
	if r != exp {
		t.Errorf("Expected %v, got %v\n", exp, r)
	}
}

func TestSolution2(t *testing.T) {
	bridgeLength := 100
	weight := 100
	truckWeights := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	r := solution(bridgeLength, weight, truckWeights)
	exp := 110
	if r != exp {
		t.Errorf("Expected %v, got %v\n", exp, r)
	}
}
