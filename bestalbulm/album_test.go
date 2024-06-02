package bestalbulm

import "testing"

func TestSolution(t *testing.T) {
	genres := []string{"classic", "pop", "classic", "classic", "pop"}
	plays := []int{500, 600, 150, 800, 2500}
	r := solution(genres, plays)
	if r[0] != 4 || r[1] != 1 || r[2] != 3 || r[3] != 0 {
		t.Error("Expected [4 1 3 0], got ", r)
	}

}
