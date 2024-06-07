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

	for i := 0; i < n; i++ {
		startx, starty, endx, endy := 0, 0, 0, 0
		_, _ = fmt.Fscanf(reader, "%d %d %d %d\n", &startx, &starty, &endx, &endy)

		cn := 0
		_, _ = fmt.Fscanf(reader, "%d\n", &cn)
		var circles [][]int
		for j := 0; j < cn; j++ {
			x, y, r := 0, 0, 0
			_, _ = fmt.Fscanf(reader, "%d %d %d\n", &x, &y, &r)
			circles = append(circles, []int{x, y, r})
		}
		r := solution(startx, starty, endx, endy, cn, circles)
		_, _ = fmt.Fprintf(writer, "%d\n", r)
	}
}
