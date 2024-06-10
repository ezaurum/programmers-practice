package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	const tt = `26 -76
65 -83
78 38
92 22
-60 -42
-27 85
42 46
-86 98
92 -47
-41 38`
	var vectors [][]int
	split := strings.Split(tt, "\n")
	for _, v := range split {
		parts := strings.Fields(v)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		vectors = append(vectors, []int{x, y})
	}
	r := solution(10, vectors)
	exp := 13.341664064126334
	if r != exp {
		t.Errorf("got %f, want %f", r, exp)
	}
}
