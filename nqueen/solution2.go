package main

func solution2(n int) int {
	ans := 0
	colchk := make([]bool, n)
	diachk := make([]bool, 2*n-1)
	rdiachk := make([]bool, 2*n-1)

	var search func(int)
	search = func(r int) {
		if r == n {
			ans++
			return
		}
		for j := 0; j < n; j++ {
			z, x, c := &colchk[j], &diachk[n-1+r-j], &rdiachk[r+j]
			if !*z && !*x && !*c {
				*z, *x, *c = true, true, true
				search(r + 1)
				*z, *x, *c = false, false, false
			}
		}
	}
	search(0)
	return ans
}
