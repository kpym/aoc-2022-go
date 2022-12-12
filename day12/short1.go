package main

import (
	"bytes"
	"os"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	nx, ny := len(lines[0]), len(lines)

	var ex, ey, sx, sy int
	mp := make([][]int, ny)
	for y, line := range lines {
		mp[y] = make([]int, nx)
		for x := range line {
			if lines[y][x] == 'S' {
				lines[y][x] = 'a'
				sx, sy = x, y
			}
			if lines[y][x] == 'E' {
				lines[y][x] = 'z'
				ex, ey = x, y
			}
			mp[y][x] = nx * ny
		}
	}

	var explore func(x, y, cost int)
	explore = func(x, y, cost int) {
		if mp[y][x] <= cost {
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

	println(mp[sy][sx])
}
