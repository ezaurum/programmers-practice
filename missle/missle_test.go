package missle

import (
	"encoding/json"
	"testing"
)

const test = `[[4,5],[4,8],[10,14],[11,13],[5,12],[3,7],[1,4]]`
const test0 = `[[620, 750], [620, 750], [620, 750]]`
const test1 = `[[0, 4], [1, 2], [1, 3], [3, 4]]`
const test2 = `[[0, 4], [0, 1], [2, 3]]`

func TestSolution(t *testing.T) {
	var input [][]int
	_ = json.Unmarshal([]byte(test), &input)
	r := solution(input)
	if r != 3 {
		t.Error("Expected 3, got ", r)
	}
}

func TestSolution0(t *testing.T) {
	var input [][]int
	_ = json.Unmarshal([]byte(test0), &input)
	r := solution(input)
	if r != 1 {
		t.Error("Expected 1, got ", r)
	}
}

func TestSolution1(t *testing.T) {
	var input [][]int
	_ = json.Unmarshal([]byte(test1), &input)
	r := solution(input)
	if r != 2 {
		t.Error("Expected 2, got ", r)
	}
}

func TestSolution2(t *testing.T) {
	var input [][]int
	_ = json.Unmarshal([]byte(test2), &input)
	r := solution(input)
	if r != 2 {
		t.Error("Expected 2, got ", r)
	}
}
