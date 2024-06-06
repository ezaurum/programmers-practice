package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acmicpc.net/problem/1002
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	var n int
	_, _ = fmt.Fscanf(reader, "%d\n", &n)

	var edges [][]int
	for i := 0; i < n; i++ {
		edge := make([]int, 6)
		_, _ = fmt.Fscanf(reader, "%d %d %d %d %d %d\n", &edge[0], &edge[1], &edge[2], &edge[3], &edge[4], &edge[5])
		edges = append(edges, edge)
	}
	r := solution(n, edges)
	for _, v := range r {
		_, _ = fmt.Fprintln(writer, v)
	}
}
