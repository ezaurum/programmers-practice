package main

import (
	"strconv"
	"strings"
	"testing"
)

const test = `3
0 0 13 40 0 37
0 0 3 0 7 4
1 1 1 1 1 5`

func TestSolution(t *testing.T) {
	var arr [][]int
	split := strings.Split(test, "\n")
	n, _ := strconv.Atoi(split[0])
	for _, i2 := range split[1:] {
		trim := strings.Trim(i2, " ")
		vs := strings.Split(trim, " ")
		var temp []int
		for _, v := range vs {
			i, err := strconv.Atoi(v)
			if err != nil {
				t.Error(err)
			}
			temp = append(temp, i)
		}
		arr = append(arr, temp)
	}
	r := solution(n, arr)
	exp := []int{2, 1, 0}
	for i, v := range r {
		if v != exp[i] {
			t.Errorf("Expected %v, got %v\n", exp, r)
		}
	}
}
