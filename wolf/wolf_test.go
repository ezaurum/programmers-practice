package wolf

import (
	"testing"
)

func TestSolution(t *testing.T) {
	info := []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1}
	edges := [][]int{{0, 1}, {1, 2}, {1, 4}, {0, 8}, {8, 7}, {9, 10}, {9, 11}, {4, 3}, {6, 5}, {4, 6}, {8, 9}}
	result := solution(info, edges)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestSolution2(t *testing.T) {
	info := []int{0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0}
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}, {3, 7}, {4, 8}, {6, 9}, {9, 10}}
	result := solution(info, edges)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestSolution3(t *testing.T) {
	info := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}, {3, 7}, {4, 8}, {6, 9}, {9, 10}}
	result := solution(info, edges)
	v := 11
	if result != v {
		t.Errorf("Expected %v, got %v", v, result)
	}
}
