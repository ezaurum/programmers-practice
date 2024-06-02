package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// https://www.acmicpc.net/problem/1710
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	text := ""
	next := true
	for next {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			next = false
		} else if err != nil {
			panic(err)
		}
		if readString == "\n" {
			next = false
		}
		if strings.HasPrefix(readString, "</body>") {
			next = false
		}
		text += readString
	}
	r := Solution(text)
	_, _ = fmt.Fprintln(writer, r)
}
