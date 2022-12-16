package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type valve struct {
	name    string
	rate    int
	to      []string
	visited bool
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	valves := make([]valve, 0, len(lines))
	shortest := make(map[string]int) // shortest["AABB"] is the shortest path from AA to BB
	for _, line := range lines {
		re := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)
		m := re.FindStringSubmatch(line)
		rate, _ := strconv.Atoi(m[2])
		to := strings.Split(m[3], ", ")
		valves = append(valves, valve{name: m[1], rate: rate, to: to})
		for _, t := range to {
			shortest[m[1]+t] = 1
		}
	}
	// fill shortest with all shortest paths
	for d, more := 2, true; more; d++ {
		more = false
		for _, v := range valves {
			for _, w := range valves {
				if _, ok := shortest[v.name+w.name]; ok {
					continue // skip if already known
				}
				// check if we can get from v to a neighbor of w
				toneighbor := false
				for _, tname := range w.to {
					if dt, ok := shortest[v.name+tname]; ok && dt < d {
						toneighbor = true
						break
					}
				}
				if toneighbor {
					more = true
					shortest[v.name+w.name] = d
				}
			}
		}
	}

	// remove usless valves (with rate 0)
	for i := 0; i < len(valves); {
		if valves[i].rate > 0 {
			i++
			continue
		}
		valves = append(valves[:i], valves[i+1:]...)
	}

	var next func(v *valve, time int) int
	next = func(v *valve, time int) int {
		time--
		if time <= 0 {
			return 0
		}
		max := 0
		for i := range valves {
			if valves[i].visited {
				continue
			}
			valves[i].visited = true
			nextmax := next(&valves[i], time-shortest[v.name+valves[i].name])
			if nextmax > max {
				max = nextmax
			}
			valves[i].visited = false
		}
		return max + v.rate*time
	}

	println(next(&valve{name: "AA"}, 31))
}
