package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	var t int
	_, _ = fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		var n, k int
		_, _ = fmt.Fscanf(reader, "%d %d\n", &n, &k)

		s, _, _ := reader.ReadLine()
		tt := strings.Split(string(s), " ")
		var times []int
		for _, s2 := range tt {
			ii, _ := strconv.Atoi(s2)
			times = append(times, ii)
		}

		var rules [][]int
		for j := 0; j < k; j++ {
			var a, b int
			_, _ = fmt.Fscanf(reader, "%d %d\n", &a, &b)
			rules = append(rules, []int{a, b})
		}
		var w int
		_, _ = fmt.Fscanf(reader, "%d\n", &w)
		r := solution(n, k, rules, w, times)
		_, _ = fmt.Fprintf(writer, "%d\n", r)
	}
}
