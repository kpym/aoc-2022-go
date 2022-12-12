package main

import (
	"bytes"
	"os"
	"sync"
)

type point struct {
	sync.Mutex
	cost int
}

func (p *point) set(cost int) (ok bool) {
	p.Lock()
	if p.cost > cost {
		p.cost = cost
		ok = true
	}
	p.Unlock()
	return
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	nx, ny := len(lines[0]), len(lines)

	var ex, ey int
	mp := make([][]point, ny)
	for y, line := range lines {
		mp[y] = make([]point, nx)
		for x := range line {
			if lines[y][x] == 'S' {
				lines[y][x] = 'a'
			}
			if lines[y][x] == 'E' {
				lines[y][x] = 'z'
				ex, ey = x, y
			}
			mp[y][x].cost = nx * ny
		}
	}

	wg := sync.WaitGroup{}
	min := point{cost: nx * ny}
	var explore func(x, y, cost int)
	explore = func(x, y, cost int) {
		defer wg.Done()
		if !mp[y][x].set(cost) {
			return
		}
		if lines[y][x] == 'a' {
			min.set(cost)
			return
		}
		if x > 0 && lines[y][x-1] >= lines[y][x]-1 {
			wg.Add(1)
			go explore(x-1, y, cost+1)
		}
		if x < nx-1 && lines[y][x+1] >= lines[y][x]-1 {
			wg.Add(1)
			go explore(x+1, y, cost+1)
		}
		if y > 0 && lines[y-1][x] >= lines[y][x]-1 {
			wg.Add(1)
			go explore(x, y-1, cost+1)
		}
		if y < ny-1 && lines[y+1][x] >= lines[y][x]-1 {
			wg.Add(1)
			go explore(x, y+1, cost+1)
		}
	}
	wg.Add(1)
	explore(ex, ey, 0)

	wg.Wait()
	println(min.cost)
}
