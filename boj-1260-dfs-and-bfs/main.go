package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://www.acmicpc.net/problem/1260
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	var n, m, start int
	_, _ = fmt.Fscanf(reader, "%d %d %d\n", &n, &m, &start)

	var edges [][]int
	for i := 0; i < m; i++ {
		edge := []int{0, 0}
		_, _ = fmt.Fscanf(reader, "%d %d\n", &edge[0], &edge[1])
		edges = append(edges, edge)
	}
	r := Solution(n, m, start, edges)
	for _, v := range r {
		_, _ = fmt.Fprintln(writer, strings.Join(v, " "))
	}
}
