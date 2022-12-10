package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	tailpos := make(map[int]bool)
	x, y := [10]int{}, [10]int{}

	for _, line := range lines {
		n, _ := strconv.Atoi(string(line[2:]))
		for i := 0; i < n; i++ {
			switch line[0] {
			case 'U':
				y[0]++
			case 'D':
				y[0]--
			case 'L':
				x[0]--
			case 'R':
				x[0]++
			}
			for k := 0; k < 9; k++ {
				if x[k] > x[k+1]+1 || x[k] < x[k+1]-1 || y[k] > y[k+1]+1 || y[k] < y[k+1]-1 {
					if x[k+1] < x[k] {
						x[k+1]++
					}
					if x[k+1] > x[k] {
						x[k+1]--
					}
					if y[k+1] < y[k] {
						y[k+1]++
					}
					if y[k+1] > y[k] {
						y[k+1]--
					}
				}
				tailpos[x[9]+y[9]<<16] = true
			}
		}
	}

	println(len(tailpos))
}
