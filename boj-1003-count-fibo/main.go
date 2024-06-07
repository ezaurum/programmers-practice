package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	var n int
	_, _ = fmt.Fscanf(reader, "%d\n", &n)

	var edges []int
	for i := 0; i < n; i++ {
		v := 0
		_, _ = fmt.Fscanf(reader, "%d\n", &v)
		edges = append(edges, v)
	}
	for _, v := range edges {
		r0, r1 := solution(v)
		_, _ = fmt.Fprintf(writer, "%d %d\n", r0, r1)
	}
}
