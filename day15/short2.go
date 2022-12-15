package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// interval represent a range [from,to] of x values that are covered by a sensor (on the current line)
// if from > to, the interval is empty
// sy is the sensor's y
type interval struct {
	from, to, sy int
}

var invls []interval

// first non-empty intervals, then empty intervals, ordered by 'from',
func less(i, j int) bool {
	if (invls[i].from <= invls[i].to) && (invls[j].from > invls[j].to) {
		return true
	}
	if (invls[i].from > invls[i].to) && (invls[j].from <= invls[j].to) {
		return false
	}
	return (invls[i].from < invls[j].from)
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	invls = make([]interval, 0, len(lines))
	for _, line := range lines {
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		dist := abs(sx-bx) + abs(sy-by)
		invls = append(invls, interval{
			from: sx - dist + abs(sy+1), // +1 because we start at y=-1
			to:   sx + dist - abs(sy+1), // +1 because we start at y=-1
			sy:   sy,
		})
	}

	max := 4000000
	for y := 0; y < max; y++ {
		// next line : some intervals shrink, some grow
		for i := 0; i < len(invls); i++ {
			if y <= invls[i].sy {
				invls[i].from--
				invls[i].to++
			} else {
				invls[i].from++
				invls[i].to--
			}
		}
		sort.Slice(invls, less)
		for x := 0; x < max; {
			// check if x is covered by an interval
			for _, s := range invls {
				if s.from > s.to { // first empty interval ?
					break
				}
				if x < s.from { // x is between two intervals or before the first interval ?
					println(x*4000000 + y)
					return
				}
				if x <= s.to {
					x = s.to + 1
				}
			}
			if x < max { // x is after the last interval ?
				println(x*4000000 + y)
				return
			}
		}
	}
}
