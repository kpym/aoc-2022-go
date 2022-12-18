package main

import (
	"bytes"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type cube [3]int

func nextto(c, d cube) bool {
	count := 0
	for i := 0; i < 3; i++ {
		count += abs(c[i] - d[i])
	}
	return count == 1
}

func neighbours(c cube, cs []cube) int {
	count := 0
	for i := range cs {
		if nextto(c, cs[i]) {
			count++
		}
	}
	return count
}

func contained(c cube, cs []cube) bool {
	for _, d := range cs {
		if c[0] == d[0] && c[1] == d[1] && c[2] == d[2] {
			return true
		}
	}
	return false
}

func surrounded(c cube, cs []cube) bool {
	return contained(cube{c[0] + 1, c[1], c[2]}, cs) && contained(cube{c[0] - 1, c[1], c[2]}, cs) && contained(cube{c[0], c[1] + 1, c[2]}, cs) && contained(cube{c[0], c[1] - 1, c[2]}, cs) && contained(cube{c[0], c[1], c[2] + 1}, cs) && contained(cube{c[0], c[1], c[2] - 1}, cs)
}

func istrapped(c cube, cs []cube) bool {
	has := [6]bool{}
	for i := range cs {
		has[0] = has[0] || (c[0] == cs[i][0] && c[1] == cs[i][1] && c[2] < cs[i][2])
		has[1] = has[1] || (c[0] == cs[i][0] && c[1] == cs[i][1] && c[2] > cs[i][2])
		has[2] = has[2] || (c[0] == cs[i][0] && c[1] < cs[i][1] && c[2] == cs[i][2])
		has[3] = has[3] || (c[0] == cs[i][0] && c[1] > cs[i][1] && c[2] == cs[i][2])
		has[4] = has[4] || (c[0] < cs[i][0] && c[1] == cs[i][1] && c[2] == cs[i][2])
		has[5] = has[5] || (c[0] > cs[i][0] && c[1] == cs[i][1] && c[2] == cs[i][2])
	}
	return has[0] && has[1] && has[2] && has[3] && has[4] && has[5]
}

func main() {
	data, _ := os.ReadFile(os.Args[1])

	var (
		cubes       []cube
		c, min, max cube
		i           int
	)
	min = cube{1 << 30, 1 << 30, 1 << 30}
	for _, v := range bytes.FieldsFunc(data, func(r rune) bool { return r < '0' || r > '9' }) {
		c[i], _ = strconv.Atoi(string(v))
		if min[i] > c[i] {
			min[i] = c[i]
		}
		if max[i] < c[i] {
			max[i] = c[i]
		}
		i++
		if i == 3 {
			cubes = append(cubes, c)
			i = 0
		}
	}

	area := 0
	for i := 0; i < len(cubes); i++ {
		area += 6 - 2*neighbours(cubes[i], cubes[:i])
	}

	var trapped []cube
	for i := min[0]; i < max[0]; i++ {
		for j := min[1]; j < max[1]; j++ {
			for k := min[2]; k < max[2]; k++ {
				c := cube{i, j, k}
				if !contained(c, cubes) && istrapped(c, cubes) {
					trapped = append(trapped, c)
				}
			}
		}
	}
	all := append(cubes, trapped...)
	n := len(cubes)
check:
	for {
		for i := 0; i < len(trapped); i++ {
			if !surrounded(trapped[i], all) {
				trapped = append(trapped[:i], trapped[i+1:]...)
				all = append(all[:n+i], all[n+i+1:]...)
				continue check
			}
		}
		break
	}

	trappedarea := 0
	for i := 0; i < len(trapped); i++ {
		trappedarea += 6 - 2*neighbours(trapped[i], trapped[:i])
	}

	println(area - trappedarea)
}
