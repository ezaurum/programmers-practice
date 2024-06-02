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

	var n, k int
	_, _ = fmt.Fscanf(reader, "%d %d\n", &n, &k)
	_, _ = fmt.Fprint(writer, Solution(n, k))
}
