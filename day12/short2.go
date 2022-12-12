package main

import (
	"bytes"
	"os"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	nx := len(lines[0])
	ny := len(lines)

	mp := make([][]int, ny)
	for y := range mp {
		mp[y] = make([]int, nx)
	}

	find := func(s, t byte) (x, y int) {
		for y, line := range lines {
			for x, c := range line {
				if c == s {
					lines[y][x] = t
					return x, y
				}
			}
		}
		return -1, -1
	}
	ex, ey := find('E', 'z')
	find('S', 'a')

	var explore func(x, y, cost int)
	explore = func(x, y, cost int) {
		if (x == ex && y == ey && cost > 0) || (mp[y][x] > 0 && mp[y][x] <= cost) {
			return
		}
		mp[y][x] = cost
		if x > 0 && lines[y][x-1] >= lines[y][x]-1 {
			explore(x-1, y, cost+1)
		}
		if x < nx-1 && lines[y][x+1] >= lines[y][x]-1 {
			explore(x+1, y, cost+1)
		}
		if y > 0 && lines[y-1][x] >= lines[y][x]-1 {
			explore(x, y-1, cost+1)
		}
		if y < ny-1 && lines[y+1][x] >= lines[y][x]-1 {
			explore(x, y+1, cost+1)
		}
	}
	explore(ex, ey, 0)

	min := nx * ny
	for y, line := range lines {
		for x, c := range line {
			if c == 'a' {
				if mp[y][x] > 0 && mp[y][x] < min {
					min = mp[y][x]
				}
			}
		}
	}

	println(min)
}
