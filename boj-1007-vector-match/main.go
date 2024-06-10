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

		vectors := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			parts = strings.Fields(scanner.Text())
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			vectors[i] = []int{x, y}
		}

		fmt.Println(solution(n, vectors))
	}
}
