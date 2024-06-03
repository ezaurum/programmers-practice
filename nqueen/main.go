package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	var n int
	_, _ = fmt.Fscanf(reader, "%d", &n)
	_, _ = fmt.Fprint(writer, solution(n))
}

// solution https://www.acmicpc.net/problem/9663
func solution(n int) int {
	// 퀸은 같은 행, 열, 대각선에 다른 퀸이 있으면 안된다.
	// 따라서 퀸은 한 행에 하나씩만 놓을 수 있다. = 행은 고정하고, 열에서 어떤 위치에 있는지만 체크하면 된다.
	columns := make([]int, n)
	return nqueen(-1, columns)
}

func promising(y int, columns []int) bool {
	for i := 0; i < y; i++ {
		if columns[i] == columns[y] {
			return false
		}
		if math.Abs(float64(columns[y]-columns[i])) == float64(y-i) {
			return false
		}
	}
	return true
}

func nqueen(y int, columns []int) int {
	if !promising(y, columns) {
		return 0
	}
	size := len(columns)
	if y == size-1 {
		return 1
	}
	count := 0
	for x := 0; x < size; x++ {
		nextRow := y + 1
		columns[nextRow] = x
		count += nqueen(nextRow, columns)
	}
	return count
}
