package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	var arr []int
	var n int
	_, _ = fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		var k int
		_, _ = fmt.Fscanf(reader, "%d\n", &k)
		arr = append(arr, k)
	}
	_, _ = fmt.Fprint(writer, strings.Join(solution(arr), "\n"))
}
