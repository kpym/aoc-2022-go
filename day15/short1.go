package main

import (
	"fmt"
	"os"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	y := 2000000
	ok := struct{}{}
	inrange := make(map[int]struct{}) // B and # positions on line y
	beacons := make(map[int]struct{}) // B positions on line y
	for _, line := range lines {
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if by == y {
			beacons[bx] = ok // beacon on (bx,y)
		}
		d := abs(sx-bx) + abs(sy-by) - abs(sy-y)
		for x := sx - d; x <= sx+d; x++ {
			inrange[x] = ok // no other beacon on (x,y)
		}
	}

	println(len(inrange) - len(beacons))
}
