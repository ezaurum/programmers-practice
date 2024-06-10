package main

// https://www.acmicpc.net/problem/1006
func solution(n, w int, enemies []int) int {
	m := make([][16]int, n)
	numPairs := greedyPairs(0, n, w, enemies, 0, m)
	answer := numPairs + (n-numPairs)*2
	return answer
}

func greedyPairs(i, n, w int, enemies []int, shadow int, m [][16]int) int {
	// fast returns
	if i == n {
		return 0
	}
	if m[i][shadow] != 0 {
		return m[i][shadow] - 1
	}

	var (
		r  = (i + 1) % n // right
		b  = i + n       // bottom
		br = r + n       // bottom right
	)

	nextShadow := func(leftBits, rightBits int) int {
		if i == 0 {
			leftBits <<= 2
		} else {
			leftBits = shadow & 12
		}
		return leftBits | rightBits
	}

	// assume no pair
	x := greedyPairs(i+1, n, w, enemies, nextShadow(0, 0), m)

	// assume left pair
	if shadow&3 == 0 && enemies[i]+enemies[b] <= w {
		here := 1
		next := greedyPairs(i+1, n, w, enemies, nextShadow(3, 0), m)
		x = max(x, here+next)
	}

	both := 0

	// assume top pair
	if shadow&1 == 0 && i != r && enemies[i]+enemies[r] <= w {
		if !(i == n-1 && shadow&4 != 0) {
			here := 1
			next := greedyPairs(i+1, n, w, enemies, nextShadow(1, 1), m)
			x = max(x, here+next)
			both++
		}
	}

	// assume bottom pair
	if shadow&2 == 0 && b != br && enemies[b]+enemies[br] <= w {
		if !(i == n-1 && shadow&8 != 0) {
			here := 1
			next := greedyPairs(i+1, n, w, enemies, nextShadow(2, 2), m)
			x = max(x, here+next)
			both++
		}
	}

	// assume both top & bottom pairs
	if both == 2 {
		here := 2
		next := greedyPairs(i+1, n, w, enemies, nextShadow(3, 3), m)
		x = max(x, here+next)
	}

	// cache it
	m[i][shadow] = x + 1
	return x
}
