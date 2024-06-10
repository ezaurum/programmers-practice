package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for t > 0 {
		t--
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])

		first := make([]int, n)
		second := make([]int, n)

		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		for i := 0; i < n; i++ {
			first[i], _ = strconv.Atoi(parts[i])
		}

		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		for i := 0; i < n; i++ {
			second[i], _ = strconv.Atoi(parts[i])
		}

		fmt.Println(solution(n, w, first, second))
	}
}
