package missle

import (
	"encoding/json"
	"testing"
)

const test = `[[0,1,1],[0,2,2],[1,2,5],[1,3,1],[2,3,8]]`

func TestSolution(t *testing.T) {
	var input [][]int
	_ = json.Unmarshal([]byte(test), &input)
	r := solution(4, input)
	if r != 4 {
		t.Error("Expected 4, got ", r)
	}
}
