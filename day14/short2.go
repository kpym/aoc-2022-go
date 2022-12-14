package main

import (
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func points(d string) (pts []point) {
	t := strings.FieldsFunc(d, func(r rune) bool {
		return (r < '0' || r > '9')
	})
	for i := 0; i < len(t); i += 2 {
		x, _ := strconv.Atoi(t[i])
		y, _ := strconv.Atoi(t[i+1])
		pts = append(pts, point{x, y})
	}
	return
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	// get the drawing instructions
	pts := make([][]point, len(lines))
	for i, l := range lines {
		pts[i] = points(l)
	}
	// get coordinates range
	minx, maxx, maxy := 1<<31, 0, 0
	for _, pp := range pts {
		for _, p := range pp {
			if p.x < minx {
				minx = p.x
			}
			if p.x > maxx {
				maxx = p.x
			}
			if p.y > maxy {
				maxy = p.y
			}
		}
	}
	// delimit the occupied array
	dy := maxy + 3
	if minx > 500-dy-1 {
		minx = 500 - dy - 1
	}
	if maxx < 500+dy+1 {
		maxx = 500 + dy + 1
	}
	dx := maxx - minx + 1
	// normalize the coordinates
	for i := range pts {
		for j := range pts[i] {
			pts[i][j].x -= minx
		}
	}
	start := point{500 - minx, 0}

	// create the "occupied" matrix
	o := make([][]bool, dy)
	for i := range o {
		o[i] = make([]bool, dx)
	}
	// fill the bottom line
	for x := 0; x < dx; x++ {
		o[dy-1][x] = true
	}

	// occupyed by the rocks
	for _, pp := range pts {
		prev := pp[0]
		for j := 1; j < len(pp); j++ {
			p := pp[j]
			if p.x == prev.x {
				miny, maxy := prev.y, p.y
				if miny > maxy {
					miny, maxy = maxy, miny
				}
				for y := miny; y <= maxy; y++ {
					o[y][p.x] = true
				}
			} else {
				minx, maxx := prev.x, p.x
				if minx > maxx {
					minx, maxx = maxx, minx
				}
				for x := minx; x <= maxx; x++ {
					o[p.y][x] = true
				}
			}
			prev = p
		}
	}

	// start dropping the sands
	sands := 0
	for { // next sand
		s := start
		if o[s.y][s.x] { // the entry is occupied
			println(sands)
			return
		}
		for { // next move
			switch {
			case !o[s.y+1][s.x]:
				s.y++
				continue
			case !o[s.y+1][s.x-1]:
				s.y++
				s.x--
				continue
			case !o[s.y+1][s.x+1]:
				s.y++
				s.x++
				continue
			}
			o[s.y][s.x] = true // stop moving
			sands++
			break
		}
	}
}
