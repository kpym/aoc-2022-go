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
	hx, hy, tx, ty := 0, 0, 0, 0

	for _, line := range lines {
		n, _ := strconv.Atoi(string(line[2:]))
		for i := 0; i < n; i++ {
			switch line[0] {
			case 'U':
				hy++
			case 'D':
				hy--
			case 'L':
				hx--
			case 'R':
				hx++
			}
			if tx < hx-1 || tx > hx+1 || ty < hy-1 || ty > hy+1 {
				if tx < hx {
					tx++
				}
				if tx > hx {
					tx--
				}
				if ty < hy {
					ty++
				}
				if ty > hy {
					ty--
				}
			}
			tailpos[tx+ty<<16] = true
		}
	}

	println(len(tailpos))
}
