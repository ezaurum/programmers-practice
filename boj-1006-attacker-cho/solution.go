package main

// https://www.acmicpc.net/problem/1006
func solution(n, w int, first, second []int) int {
	/*

	   한 특수소대는 침투한 구역 외에, 인접한 한 구역 더 침투할 수 있다. (같은 경계를 공유하고 있으면 인접 하다고 한다. 위 그림에서 1구역은 2, 8, 9 구역과 서로 인접한 상태다.) 즉, 한 특수소대는 한 개 혹은 두 개의 구역을 커버할 수 있다.
	   특수소대끼리는 아군인지 적인지 구분을 못 하기 때문에, 각 구역은 하나의 소대로만 커버해야 한다.
	   한 특수소대가 커버하는 구역의 적들의 합은 특수소대원 수 W 보다 작거나 같아야 한다.

	*/
	// 일단 선형으로 풀어보자
	arr := [][]int{first, second}

	if n == 1 {
		return check1or2(arr[0][1], arr[1][1], w)
	}

	var dp [][]int
	dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	solve := func() {
		for i := 2; i < n; i++ {
			up := check1or2(arr[0][i], arr[0][i-1], w)
			down := check1or2(arr[1][i], arr[1][i-1], w)
			ver := check1or2(arr[0][i], arr[1][i], w)

			dp[i][0] = min(dp[i-1][0]+ver, dp[i-2][0]+up+down, dp[i-1][1]+up, dp[i-1][2]+down)
			dp[i][1] = min(dp[i-1][2]+down, dp[i-1][0]+1)
			dp[i][2] = min(dp[i-1][1]+up, dp[i-1][0]+1)
		}
	}

	// 초기화는 최대한으로
	over := len(first) + len(second) + 1
	result := over

	// 모두 연결 안 되어있을 때
	dp[1][0] = check1or2(arr[0][1], arr[1][1], w)
	dp[1][1] = 1
	dp[1][2] = 1
	solve()
	result = min(result, dp[n-1][0])

	// 위가 연결되어있을 때
	if arr[0][1]+arr[1][1] <= w {
		dp[1][0] = 2
		dp[1][1] = 1
		dp[1][2] = 1
		t := arr[0][1]
		arr[0][1] = over
		solve()
		result = min(result, dp[n-1][1])
		arr[0][1] = t
	}

	// 아래가 연결되어있을 때
	if arr[1][1]+arr[1][n-1] <= w {
		dp[1][0] = 2
		dp[1][1] = 1
		dp[1][2] = 1
		t := arr[1][1]
		arr[1][1] = over
		solve()
		result = min(result, dp[n-1][2])
		arr[1][1] = t
	}

	// 위 아래가 연결되어있을 때

	if arr[0][1]+arr[0][n-1] <= w && arr[1][1]+arr[1][n-1] <= w {
		tUp := arr[0][1]
		tDown := arr[1][1]
		arr[0][1] = over
		arr[1][1] = over
		dp[1][0] = 2
		dp[1][1] = 1
		dp[1][2] = 1
		solve()
		result = min(result, dp[n-1][0])
		arr[0][1] = tUp
		arr[1][1] = tDown
	}

	return result
}

func check1or2(a, b, w int) int {
	if a+b <= w {
		return 1
	}
	return 2

}
