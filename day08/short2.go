package main

import (
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	l, r, max := len(lines), len(lines[0]), 0

	var k int // temp variable used in loops
	for i := 1; i < r-1; i++ {
		for j := 1; j < l-1; j++ {
			score := 1

			// left
			for k = j - 1; k > 0 && lines[i][k] < lines[i][j]; k-- {
			}
			score *= j - k
			// right
			for k = j + 1; k < l-1 && lines[i][k] < lines[i][j]; k++ {
			}
			score *= k - j
			// up
			for k = i - 1; k > 0 && lines[k][j] < lines[i][j]; k-- {
			}
			score *= i - k
			// down
			for k = i + 1; k < r-1 && lines[k][j] < lines[i][j]; k++ {
			}

			score *= k - i
			if score > max {
				max = score
			}
		}
	}

	println(max)
}
