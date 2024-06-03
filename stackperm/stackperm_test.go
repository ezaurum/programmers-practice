package main

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	arr := []int{4, 3, 6, 8, 7, 5, 2, 1}
	r := solution(arr)
	exp := `+
+
+
+
-
-
+
+
-
+
+
-
-
-
-
-
`
	join := strings.Join(r, "\n")
	if join != strings.Trim(exp, "\n") {
		t.Errorf("Expected %v, got %v\n", exp, join)
	}
}

func TestSolution0(t *testing.T) {
	arr := []int{1, 2, 5, 3, 4}
	r := solution(arr)
	exp := `NO`
	join := strings.Join(r, "\n")
	if join != strings.Trim(exp, "\n") {
		t.Errorf("Expected %v, got %v\n", exp, join)
	}
}
