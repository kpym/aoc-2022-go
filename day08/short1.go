package main

import (
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	l, r, count := len(lines), len(lines[0]), 0

	var k int
	for i := 0; i < r; i++ {
		for j := 0; j < l; j++ {
			// left
			for k = j - 1; k >= 0 && lines[i][k] < lines[i][j]; k-- {
			}
			if k < 0 {
				count++
				continue
			}
			// right
			for k = j + 1; k < r && lines[i][k] < lines[i][j]; k++ {
			}
			if k == r {
				count++
				continue
			}
			// top
			for k = i - 1; k >= 0 && lines[k][j] < lines[i][j]; k-- {
			}
			if k < 0 {
				count++
				continue
			}
			// bottom
			for k = i + 1; k < l && lines[k][j] < lines[i][j]; k++ {
			}
			if k == l {
				count++
				continue
			}
		}
	}

	println(count)
}
